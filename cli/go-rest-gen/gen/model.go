package gen

import (
	"fmt"
	"strings"
)

func GenerateModel(packageName string, config Config) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, "\t\"time\"\n")
	buf = append(buf, "\t\"go.mongodb.org/mongo-driver/bson/primitive\"\n")

	// buf = append(buf, fmt.Sprintf("\t\"app/services/%s\"", packageName))
	for _, v := range config.Import.Model {
		buf = append(buf, fmt.Sprintf("\t\"%s\"", v))
	}

	buf = append(buf, ")\n")

	buf = append(buf, fmt.Sprintf("const CollectionName = \"%s\"\n", config.Collection))

	for _, v := range config.Models {
		buf = append(buf, fmt.Sprintf("type %s struct {", v.Name))

		if v.Name == MainModelName {
			// ObjectID
			buf = append(buf,
				"\tId *primitive.ObjectID `bson:\"_id,omitempty\" json:\"id,omitempty\"`",
			)
		}

		for _, attr := range v.Attributes {
			var validate = ""
			if attr.Required {
				validate = " validate:\"required\""
			}

			if !attr.AllowsEmpty {
				buf = append(buf, fmt.Sprintf(
					"\t%s %s `bson:\"%s,omitempty\" json:\"%s,omitempty\"%s`",
					strings.Title(attr.Name), attr.Type, attr.Name, attr.Name, validate),
				)
			} else {
				buf = append(buf, fmt.Sprintf(
					"\t%s %s `bson:\"%s\" json:\"%s\"%s`",
					strings.Title(attr.Name), attr.Type, attr.Name, attr.Name, validate),
				)
			}
		}

		// Timestamp
		if v.Name == MainModelName {
			buf = append(buf,
				"\tCreatedAt *time.Time `bson:\"createdAt,omitempty\" json:\"createdAt,omitempty\"`",
			)

			buf = append(buf,
				"\tUpdatedAt *time.Time `bson:\"updatedAt,omitempty\" json:\"updatedAt,omitempty\"`",
			)
		}

		buf = append(buf, "}\n")
	}

	return strings.Join(buf, "\n")
}

func GenerateModelService(packageName string) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))
	buf = append(buf, "import (")
	buf = append(buf, fmt.Sprintf("\t\"app/generated/%s\"", packageName))
	buf = append(buf, ")\n")
	buf = append(buf, "type Model struct {")
	buf = append(buf, fmt.Sprintf("\t%s.Model `bson:\",inline\"`", packageName))
	buf = append(buf, "}\n")

	return strings.Join(buf, "\n")
}
