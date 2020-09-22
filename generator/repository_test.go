package generator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleRepository = `
package test

import (
	"github.com/ez-connect/go-rest/rest"
	"go.mongodb.org/mongo-driver/bson"

	"app/shared/driver"
)

type Repository struct {
	rest.RepositoryBase
	rest.RepositoryInterface
}
`

func TestGenerateRepository(t *testing.T) {
	v := GenerateRepository(
		"test", "testCollection",
	)

	assert.Equal(t, true, strings.Contains(v, sampleRepository))
}
