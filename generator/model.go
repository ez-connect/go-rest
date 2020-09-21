package generator

import (
	"fmt"
	"strings"
)

func GenerateModel(packageName string, config ModelConfig) string {
	buf := []string{}
	buf = append(buf, "\n")
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, "\t\"time\"\n")
	buf = append(buf, "\t\"go.mongodb.org/mongo-driver/bson\"")
	buf = append(buf, "\t\"go.mongodb.org/mongo-driver/bson/primitive\"")
	buf = append(buf, ")\n")

	buf = append(buf, "type Model struct {")
	for _, v := range config.Attributes {
		buf = append(buf, fmt.Sprintf(
			"\t%s %s `bson:\"%s\" json:\"%s\"`",
			strings.Title(v.Name), v.Type, v.Name, v.Name),
		)
	}
	buf = append(buf, "}")

	return strings.Join(buf, "\n")
}
