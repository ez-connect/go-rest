package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type xxx struct {
	id  *primitive.ObjectID
	abc string
}

func TestUnmarshal(t *testing.T) {
	query := Unmarshal(`{
		"$and": [
			{"abc": "xyz"},
			{"id": "5f6b0b42d59a0aa2d1906fd2"},
			{
				"$or": [
					{"abc": "xyz"},
					{"id": "5f6b0b42d59a0aa2d1906fd2"}
				]
			}
		]
	}`, xxx{})
	assert.NotNil(t, query)
}
