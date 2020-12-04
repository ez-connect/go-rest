package gen

import (
	"fmt"
	"strings"
)

func GenerateProtobuf(packageName string, config Config) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))
	buf = append(buf, "syntax = \"proto3\";\n")
	buf = append(buf, fmt.Sprintf("option go_package = \"app/generate/%s\";\n", packageName))

	/// Messages
	for _, v := range config.Models {
		buf = append(buf, fmt.Sprintf("message %s {", v.Name))
		buf = append(buf, "}\n")
	}

	buf = append(buf, fmt.Sprintf("service %s {", strings.Title(packageName)))
	///
	buf = append(buf, "}\n")

	return strings.Join(buf, "\n")
}
