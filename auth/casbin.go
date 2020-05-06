package auth

import (
	"errors"
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
)

type CasbinConfig struct {
	Model  string `yaml:"model"`
	Policy string `yaml:"policy"`
}

///////////////////////////////////////////////////////////////////

var enforcer *casbin.Enforcer

// InitCasbin creates an enforcer via file.
// See `casbin.NewEnforcer` for more info.
func InitCasbin(config CasbinConfig) error {
	model := config.Model
	policy := config.Policy
	fmt.Println("Init casbin: model =", model, "policy =", policy)

	if model == "" || policy == "" {
		return errors.New("Require both model & policy file name")
	}

	var err error
	enforcer, err = casbin.NewEnforcer(model, policy)
	return err
}

// HasPermission return true if a subject has permission on a request
func HasPermission(c echo.Context, subject string) bool {
	path := c.Path()
	method := c.Request().Method
	ok, err := enforcer.Enforce(subject, path, method)
	if err != nil {
		fmt.Println(err)
	}

	return ok
}
