package generator

import (
	"fmt"
	"strings"
)

func GenerateModel(packageName string, config ModelConfig) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, "\t\"time\"\n")
	buf = append(buf, "\t\"go.mongodb.org/mongo-driver/bson/primitive\"")
	buf = append(buf, ")\n")

	buf = append(buf, fmt.Sprintf("const CollectionName = \"%ss\"\n", config.Name))

	buf = append(buf, "type Model struct {")

	// ObjectID
	buf = append(buf,
		"\tId *primitive.ObjectID `bson:\"_id,omitempty\" json:\"id,omitempty\"`",
	)
	for _, v := range config.Attributes {
		buf = append(buf, fmt.Sprintf(
			"\t%s %s `bson:\"%s,omitempty\" json:\"%s,omitempty\"`",
			strings.Title(v.Name), v.Type, v.Name, v.Name),
		)
	}

	// Timestamp
	buf = append(buf,
		"\tCreatedAt *time.Time `bson:\"createdAt,omitempty\" json:\"createdAt,omitempty\"`",
	)

	buf = append(buf,
		"\tUpdatedAt *time.Time `bson:\"updatedAt,omitempty\" json:\"updatedAt,omitempty\"`",
	)

	buf = append(buf, "}\n")

	return strings.Join(buf, "\n")
}

func GenerateModelExt(packageName string) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, fmt.Sprintf("\t\"app/generated/%s\"", packageName))
	buf = append(buf, ")\n")

	buf = append(buf, "type Model struct {")
	buf = append(buf, fmt.Sprintf("\t%s.Model", packageName))
	buf = append(buf, "}\n")

	return strings.Join(buf, "\n")
}
