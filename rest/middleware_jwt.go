package rest

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

///////////////////////////////////////////////////////////////////

type JWTConfig struct {
	// Algorithm: HS256, HS384, HS512, RS256, RS384, RS512, ES256, ES384, ES512, None
	SigningMethod string `yaml:"signingMethod"`
	// Secrect key
	SigningKey string `yaml:"signingKey"`
	// How many ms that token will be expired
	Expire time.Duration `yaml:"expire"`
	// How many ms that refresh token will be expired
	RefreshExpire time.Duration `yaml:"refreshExpire"`
}

func (c *JWTConfig) IsValid() bool {
	if c.SigningMethod == "" || c.SigningKey == "" || c.Expire == 0 {
		return false
	}

	return true
}

///////////////////////////////////////////////////////////////////

var jwtConfig JWTConfig

func GetJWTConfig() JWTConfig {
	return jwtConfig
}

// Must be call before use JWT
func InitJWTMiddleware(config JWTConfig) error {
	fmt.Println("Init JWT: expire =", config.Expire, "refreshExpire =", config.RefreshExpire)
	if !config.IsValid() {
		return errors.New("Invalid config")
	}

	jwtConfig = config
	return nil
}

// JWT signed string
func GetJWTSignedString(data jwt.MapClaims) (string, error) {
	return getJWTSignedString(data, jwtConfig.Expire)
}

// JWT signed string with long expire
func GetRefreshJWTSignedString(data jwt.MapClaims) (string, error) {
	return getJWTSignedString(data, jwtConfig.RefreshExpire)
}

// Decode token
func GetJWTClaims(c echo.Context) jwt.MapClaims {
	data := c.Get(middleware.DefaultJWTConfig.ContextKey)
	if data == nil {
		return nil
	}

	token := data.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	return claims
}

func JWTWithDefault(skipper middleware.Skipper) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: jwtConfig.SigningMethod,
		SigningKey:    []byte(jwtConfig.SigningKey),
		Skipper:       skipper,
	})
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, jwtKeyFunc)
}

///////////////////////////////////////////////////////////////////

type MiddlewareHandler func(next echo.HandlerFunc, c echo.Context) error

// Extracts token from the request header.
func GetJWTFromHeader(c echo.Context) (string, error) {
	header := c.Request().Header.Get(echo.HeaderAuthorization)
	if header == "" {
		return "", nil
	}
	authScheme := middleware.DefaultJWTConfig.AuthScheme
	l := len(authScheme)
	if len(header) > l+1 && header[:l] == authScheme {
		return header[l+1:], nil
	}
	return "", middleware.ErrJWTMissing
}

func JWTWithAuthHandler(handler MiddlewareHandler) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth, err := GetJWTFromHeader(c)
			if err != nil {
				return err
			} else if auth == "" {
				return handler(next, c)
			}

			/// https://github.com/labstack/echo/blob/master/middleware/jwt.go#L192
			token := new(jwt.Token)
			// Issue #647, #656
			token, err = jwt.Parse(auth, jwtKeyFunc)
			if err == nil && token.Valid {
				// Store user information from token into context.
				c.Set(middleware.DefaultJWTConfig.ContextKey, token)
				return handler(next, c)
			}

			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  "invalid or expired jwt",
				Internal: err,
			}
		}
	}
}

///////////////////////////////////////////////////////////////////

// https://github.com/labstack/echo/blob/master/middleware/jwt.go#L142
func jwtKeyFunc(t *jwt.Token) (interface{}, error) {
	if t.Method.Alg() != jwtConfig.SigningMethod {
		return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
	}

	signingKey := []byte(jwtConfig.SigningKey)
	return signingKey, nil
}

func getJWTSignedString(data jwt.MapClaims, exp time.Duration) (string, error) {
	// Create token
	token := jwt.New(jwt.GetSigningMethod(jwtConfig.SigningMethod))

	// Add default expire
	if data["exp"] == nil {
		data["exp"] = time.Now().Add(exp).Unix()
	}

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	for k, v := range data {
		claims[k] = v
	}

	// Generate encoded token
	return token.SignedString([]byte(jwtConfig.SigningKey))
}
