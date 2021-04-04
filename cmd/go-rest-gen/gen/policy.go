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
				p := r.Policy
				if p == "" {
					p = "write"
					if r.Method == "GET" {
						p = "read"
					} else if r.Method == "DELETE" {
						p = "delete"
					}
				}
				buf = append(buf, fmt.Sprintf("p, %s_%s, %s%s, %s", strings.ToLower(c.Collection), p, rs.Path, r.Path, r.Method))
			}
		}
	}
	return strings.Join(buf, "\n")
}
