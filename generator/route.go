package generator

import (
	"fmt"
	"strings"
)

func GenerateRoute(packageName, collection string, routes []RouteGroup) string {
	buf := []string{}
	buf = append(buf, "\n")
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, "\t\"github.com/ez-connect/go-rest/db\"")
	buf = append(buf, "\t\"github.com/ez-connect/go-rest/rest\"")
	buf = append(buf, "\t\"github.com/labstack/echo/v4\"\n")

	buf = append(buf, "\t\"app/shared/driver\"")
	buf = append(buf, "\t\"app/shared/util\"")
	buf = append(buf, ")\n")

	buf = append(buf, "type Router struct {")
	buf = append(buf, "\trest.RouterBase")
	buf = append(buf, "}\n")

	buf = append(buf, "func (r *Router) Init(e *echo.Echo, db db.DatabaseBase) {")
	buf = append(buf, "\th := Handler{}")
	buf = append(buf, fmt.Sprintf("\th.Init(db, driver.Collection%s", strings.Title(collection)))
	buf = append(buf, "\th.Repo.Init(db)")
	buf = append(buf, "\th.Repo.EnsureIndexs()\n")

	for i, v := range routes {
		buf = append(buf, fmt.Sprintf("\tg%v := e.Group(%s)", i, v.Path))
		for _, r := range v.Children {
			buf = append(buf, fmt.Sprintf("\tg.%s(\"%s\", h.%s)", r.Method, r.Path, r.Handler))
		}
	}

	buf = append(buf, "}\n")

	return strings.Join(buf, "\n")
}
