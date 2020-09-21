package generator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleHandler = `
package test

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
`

func TestGenerateHandler(t *testing.T) {
	v := GenerateHandler(
		"test",
		ModelConfig{
			Name: "ImAModel",
			Attributes: []Attribute{
				{
					Name: "id",
					Type: "*primitive.ObjectID",
				},
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "createdAt",
					Type: "*time.Time",
				},
			},
		},
	)

	t.Error(v)
	assert.Equal(t, true, strings.Contains(v, sampleHandler))
}
