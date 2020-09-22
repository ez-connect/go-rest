package generator

import (
	"fmt"
	"strings"
)

func GenerateRepository(packageName, collection string) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "type Repository struct {")
	buf = append(buf, "\trest.RepositoryBase")
	buf = append(buf, "\trest.RepositoryInterface")
	buf = append(buf, "}\n")

	buf = append(buf, "///////////////////////////////////////////////////////////////////\n")

	buf = append(buf, "func (r *Repository) EnsureIndexs() {")
	buf = append(buf, "}\n")

	return strings.Join(buf, "\n")
}
