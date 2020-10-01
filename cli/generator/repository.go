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
	buf = append(buf, fmt.Sprintf("\t\"app/services/%s\"", packageName))
	buf = append(buf, ")\n")

	buf = append(buf, "type Repository struct {")
	buf = append(buf, fmt.Sprintf("\t%s.Repository", packageName))
	buf = append(buf, "}\n")

	buf = append(buf, "///////////////////////////////////////////////////////////////////\n")

	buf = append(buf, "func (r *Repository) EnsureIndexs() {")
	buf = append(buf, "\tr.Driver.EnsureIndex(CollectionName, \"createdAt\", bson.M{\"createdAt\": -1}, false)")
	buf = append(buf, "\tr.Driver.EnsureIndex(CollectionName, \"updatedAt\", bson.M{\"updatedAt\": -1}, false)")

	for _, v := range config.Indexes {
		name := strings.Join(v.Fields, ".")
		fields := []string{}
		for _, f := range v.Fields {
			if v.Text {
				fields = append(fields, fmt.Sprintf("\"%s\" : \"text\"", f))
			} else {
				fields = append(fields, fmt.Sprintf("\"%s\" : 1", f))
			}
		}
		buf = append(buf, fmt.Sprintf("\tr.Driver.EnsureIndex(CollectionName, \"%s\", bson.M{%s}, %v)", name, strings.Join(fields, ","), v.Unique))
	}

	buf = append(buf, "}\n")

	return strings.Join(buf, "\n")
}

func GenerateRepositoryExt(packageName string) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, "\t\"github.com/ez-connect/go-rest/rest\"")
	buf = append(buf, ")\n")

	buf = append(buf, "type Repository struct {")
	buf = append(buf, "\trest.RepositoryBase")
	buf = append(buf, "\trest.RepositoryInterface")
	buf = append(buf, "}\n")

	return strings.Join(buf, "\n")
}
