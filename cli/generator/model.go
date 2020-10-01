package generator

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
	buf = append(buf, fmt.Sprintf("\t\"app/services/%s\"", packageName))
	buf = append(buf, ")\n")

	buf = append(buf, fmt.Sprintf("const CollectionName = \"%s\"\n", config.Model.Name))

	buf = append(buf, "type Model struct {")
	buf = append(buf, fmt.Sprintf("\t%s.Model", packageName))

	// ObjectID
	buf = append(buf,
		"\tId *primitive.ObjectID `bson:\"_id,omitempty\" json:\"id,omitempty\"`",
	)
	for _, v := range config.Model.Attributes {
		var omitempty = ""
		if v.Omitempty {
			omitempty = ",omitempty"
		}

		var validate = ""
		if v.Required {
			validate = " validate:\"required\""
		}

		buf = append(buf, fmt.Sprintf(
			"\t%s %s `bson:\"%s%s\" json:\"%s%s\"%s`",
			strings.Title(v.Name), v.Type, v.Name, omitempty, v.Name, omitempty, validate),
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

	buf = append(buf, generateEmbedModels(config.EmbedModels))

	return strings.Join(buf, "\n")
}

func GenerateModelExt(packageName string) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "type Model struct {")
	buf = append(buf, "}\n")

	return strings.Join(buf, "\n")
}

func generateEmbedModels(embedConfig []ModelConfig) string {
	buf := []string{}

	for _, config := range embedConfig {
		buf = append(buf, fmt.Sprintf("type %s struct {", config.Name))

		for _, v := range config.Attributes {
			buf = append(buf, fmt.Sprintf(
				"\t%s %s `bson:\"%s,omitempty\" json:\"%s,omitempty\"`",
				strings.Title(v.Name), v.Type, v.Name, v.Name),
			)
		}

		buf = append(buf, "}\n")
	}
	return strings.Join(buf, "\n")
}
