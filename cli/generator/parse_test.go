package generator

import (
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	pkgMap, err := parser.ParseDir(token.NewFileSet(), "C:\\Projects\\go-rest\\rest", nil, parser.AllErrors)
	assert.NoError(t, err)
	assert.NotEmpty(t, pkgMap)
	for pkg := range pkgMap {
		assert.NotNil(t, pkgMap[pkg].Files)
	}
}
