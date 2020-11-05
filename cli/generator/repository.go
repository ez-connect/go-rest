package generator

import (
	"fmt"
	"strings"
)

func GenerateRepository(packageName string, config Config) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, "\t\"go.mongodb.org/mongo-driver/bson\"\n")
	buf = append(buf, "\t\"github.com/ez-connect/go-rest/rest\"")
	// buf = append(buf, fmt.Sprintf("\t\"app/services/%s\"", packageName))
	buf = append(buf, ")\n")

	buf = append(buf, "type Repository struct {")
	buf = append(buf, "\trest.RepositoryBase")
	buf = append(buf, "\trest.RepositoryInterface")

	// buf = append(buf, fmt.Sprintf("\t%s.Repository", packageName))
	for _, v := range config.Import.Repository {
		buf = append(buf, fmt.Sprintf("\t\"%s\"", v))
	}

	buf = append(buf, "}\n")

	buf = append(buf, "///////////////////////////////////////////////////////////////////\n")

	buf = append(buf, "func (r *Repository) EnsureIndexs() {")
	buf = append(buf, "\tr.Driver.EnsureIndex(CollectionName, \"createdAt\", bson.M{\"createdAt\": -1}, false)")
	buf = append(buf, "\tr.Driver.EnsureIndex(CollectionName, \"updatedAt\", bson.M{\"updatedAt\": -1}, false)")

	// Single indexes
	for _, v := range config.Index.Singles {
		order := 1
		if v.Order == -1 {
			order = -1
		}
		buf = append(buf, fmt.Sprintf("\tr.Driver.EnsureIndex(CollectionName, \"%s\", bson.M{\"%s\": %v}, %v)", v.Field, v.Field, order, v.Unique))
	}

	// Compound indexes
	for _, v := range config.Index.Compounds {
		names := []string{}
		fields := []string{}
		for _, f := range v.Fields {
			names = append(names, f.Field)
			fields = append(fields, fmt.Sprintf("\"%s\": %v", f.Field, f.Order))
		}
		buf = append(buf, fmt.Sprintf(
			"\tr.Driver.EnsureIndex(CollectionName, \"%s\", bson.M{%s}, %v)",
			strings.Join(names, "."),
			strings.Join(fields, ", "),
			v.Unique,
		))
	}

	// Text indexes
	texts := []string{}
	for _, v := range config.Index.Texts {
		texts = append(texts, fmt.Sprintf("\"%s\": \"text\"", v))
	}

	buf = append(buf, fmt.Sprintf(
		"\tr.Driver.EnsureIndex(CollectionName, \"%s\", bson.M{%s}, false)",
		"text",
		strings.Join(texts, ", "),
	))

	buf = append(buf, "}\n")

	return strings.Join(buf, "\n")
}

func GenerateRepositoryService(packageName string) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	// buf = append(buf, "\t\"github.com/ez-connect/go-rest/rest\"")
	buf = append(buf, fmt.Sprintf("\t\"app/generated/%s\"", packageName))
	buf = append(buf, ")\n")

	buf = append(buf, "type Repository struct {")
	// buf = append(buf, "\trest.RepositoryInterface")
	buf = append(buf, fmt.Sprintf("\t%s.Repository", packageName))
	buf = append(buf, "}\n")

	return strings.Join(buf, "\n")
}
