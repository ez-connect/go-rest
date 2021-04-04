package gen

import (
	"fmt"
	"strings"
)

func GeneratePolicy(config []Config) string {
	buf := []string{}
	for _, c := range config {
		for _, rs := range c.Routes {
			for _, r := range rs.Children {
				buf = append(buf, fmt.Sprintf("p, xxx, %s%s, %s", rs.Path, r.Path, r.Method))
			}
		}
	}
	return strings.Join(buf, "\n")
}
