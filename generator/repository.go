package generator

import (
	"fmt"
	"strings"
)

func GenerateRepository(packageName, collection string) string {
	buf := []string{}
	buf = append(buf, "\n")
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, "\t\"github.com/ez-connect/go-rest/rest\"")
	buf = append(buf, "\t\"go.mongodb.org/mongo-driver/bson\"\n")

	buf = append(buf, "\t\"app/shared/driver\"")
	buf = append(buf, ")\n")

	buf = append(buf, "type Repository struct {")
	buf = append(buf, "\trest.RepositoryBase")
	buf = append(buf, "\trest.RepositoryInterface")
	buf = append(buf, "}\n")

	buf = append(buf, "///////////////////////////////////////////////////////////////////\n")

	buf = append(buf, "func (r *Repository) EnsureIndexs() {")
	// buf = append(buf, "\tc := driver.CollectionBanner")
	buf = append(buf, "}\n")

	return strings.Join(buf, "\n")
}
