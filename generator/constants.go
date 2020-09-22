package generator

import (
	"fmt"
	"strings"
)

func GenerateConstant(packageName string, collections []string) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	for _, v := range collections {
		buf = append(buf, fmt.Sprintf("const Collection%s = %s", v, v))
	}

	return strings.Join(buf, "\n")
}
