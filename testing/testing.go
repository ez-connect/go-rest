package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/ez-connect/go-rest/rest"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type RequestConfig struct {
	// Default GET
	Method string
	// Default "/"
	Target string

	Path           string
	PathParamName  string
	PathParamValue string

	HeaderAuthorization string
	Body                interface{}
}

func TestFields(t *testing.T, doc interface{}) {
	v := reflect.ValueOf(doc)
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.String {
			assert.NotEqual(t, "", f.String())
		}

		if f.Type().String() == "time.Time" {
			date := f.Interface()
			var now time.Time
			assert.Equal(t, fmt.Sprintf("%v", now), fmt.Sprintf("%v", date))
		}
	}
}

func MakeRequest(e *echo.Echo, config *RequestConfig) (echo.Context, *httptest.ResponseRecorder) {
	var method string
	if config != nil && config.Method != "" {
		method = config.Method
	} else {
		method = http.MethodGet
	}

	var target string
	if config != nil && config.Target != "" {
		target = config.Target
	} else {
		target = "/"
	}

	var body io.Reader
	if config != nil && config.Body != nil {
		data, _ := json.Marshal(config.Body)
		body = bytes.NewReader(data)
	}

	req := httptest.NewRequest(
		method,
		target,
		body,
	)

	if config != nil && config.HeaderAuthorization != "" {
		req.Header.Set(echo.HeaderAuthorization, config.HeaderAuthorization)
	}

	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	if config != nil && config.Path != "" {
		c.SetPath(config.Path)
		if config.PathParamName != "" {
			c.SetParamNames(config.PathParamName)
		}
		if config.PathParamValue != "" {
			c.SetParamValues(config.PathParamValue)
		}
	}

	if config != nil && config.HeaderAuthorization != "" {
		middlewareFunc := rest.JWTWithDefault(nil)
		handlerFunc := middlewareFunc(func(c echo.Context) error {
			return nil
		})
		handlerFunc(c)
	}

	return c, res
}

func PrintResponse(t *testing.T, res *httptest.ResponseRecorder) {
	data := &bytes.Buffer{}
	json.Indent(data, res.Body.Bytes(), "", "  ")
	t.Error(data)
}
