package gen

import (
	"fmt"
	"strings"
)

func GeneratePolicy(config []Config) string {
	buf := []string{}
	for _, c := range config {
		col := strings.ToLower(c.Collection)
		for _, rs := range c.Routes {
			g := col
			if rs.Name != "" {
				g += "_" + rs.Name
			}
			for _, r := range rs.Children {
				p := r.Policy
				if p == "" {
					p = "write"
					if r.Method == "GET" {
						p = "read"
					} else if r.Method == "DELETE" {
						p = "delete"
					}
				}
				buf = append(buf, fmt.Sprintf("p, %s_%s, %s%s, %s", g, p, rs.Path, r.Path, r.Method))
			}
		}
		buf = append(buf, "")
	}
	return strings.Join(buf, "\n")
}
