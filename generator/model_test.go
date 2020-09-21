package generator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleModel = `
package test

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"

)
`

func TestGenerateModel(t *testing.T) {
	v := GenerateModel(
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

	assert.Equal(t, true, strings.Contains(v, sampleModel))
}
