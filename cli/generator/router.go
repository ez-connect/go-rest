package generator

import (
	"fmt"
	"strings"
)

func GenerateRoutes(packageName, collection string, routes []RouteGroup) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, "\t\"github.com/ez-connect/go-rest/db\"")
	buf = append(buf, "\t\"github.com/ez-connect/go-rest/rest\"")
	buf = append(buf, "\t\"github.com/labstack/echo/v4\"")
	buf = append(buf, ")\n")

	buf = append(buf, "type Router struct {")
	buf = append(buf, "\trest.RouterBase")
	buf = append(buf, "}\n")

	buf = append(buf, "func (r *Router) Init(e *echo.Echo, db db.DatabaseBase) {")
	buf = append(buf, "\th := Handler{}")
	buf = append(buf, fmt.Sprintf("\th.Init(db, Collection%s)", strings.Title(collection)))
	buf = append(buf, "\th.Repo.Init(db)")
	buf = append(buf, "\th.Repo.EnsureIndexs()\n")

	for i, v := range routes {
		buf = append(buf, fmt.Sprintf("\tg%v := e.Group(\"%s\")", i, v.Path))
		for _, r := range v.Children {
			buf = append(buf,
				fmt.Sprintf("\tg%v.%s(\"%s\", h.%s%s)", i, r.Method, r.Path, r.Handler, strings.Title(packageName)),
			)
		}
	}

	buf = append(buf, "}\n")

	return strings.Join(buf, "\n")
}

func GenerateRoutesExt(packageName string) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, fmt.Sprintf("\t\"app/generated/%s\"", packageName))
	buf = append(buf, ")\n")

	buf = append(buf, "type Router struct {")
	buf = append(buf, fmt.Sprintf("\t%s.Router", packageName))
	buf = append(buf, "}\n")

	return strings.Join(buf, "\n")
}
