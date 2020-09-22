package generator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleHandler = `
package test

import (
	"github.com/ez-connect/go-rest/rest"
	"github.com/labstack/echo/v4"
)

type Handler struct {
`

func TestGenerateHandler(t *testing.T) {
	v := GenerateHandler(
		"test",
	)

	assert.Equal(t, true, strings.Contains(v, sampleHandler))
}
