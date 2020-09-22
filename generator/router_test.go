package generator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleRoutes = `
package test

import (
	"github.com/ez-connect/go-rest/db"
	"github.com/ez-connect/go-rest/rest"
	"github.com/labstack/echo/v4"

	"app/shared/driver"
	"app/shared/util"
)

type Router struct {
	rest.RouterBase
}

func (r *Router) Init(e *echo.Echo, db db.DatabaseBase) {
`

func TestGenerateRoutes(t *testing.T) {
	v := GenerateRoutes(
		"test",
		"testModel",
		[]RouteGroup{
			{
				Path: "hello",
				Children: []RouteConfig{
					{
						Method:  "GET",
						Handler: "Find",
					},
				},
			},
		},
	)

	assert.Equal(t, true, strings.Contains(v, sampleRoutes))
}
