package gen

import (
	"fmt"
	"strings"
)

func GenerateConstants(configs []Config) string {
	buf := []string{}
	buf = append(buf, "package generated\n")
	buf = append(buf, "const (")

	for _, v := range configs {
		buf = append(buf, fmt.Sprintf("\tCollection%s = \"%s\"", strings.Title(v.Collection), v.Collection))
	}

	buf = append(buf, ")\n")

	return strings.Join(buf, "\n")
}
