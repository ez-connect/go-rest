package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	assert.NotEmpty(t, v)
}