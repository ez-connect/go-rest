package gen

import (
	"fmt"
	"strings"
)

var find = `func (h *Handler) Find(c echo.Context) error {
	docs := []Model{}
	return h.HandlerBase.Find(c, nil, &docs)
}
`

var insert = `func (h *Handler) Insert(c echo.Context) error {
	doc := Model{
		%s.Model{
			CreatedAt: core.Now(),
			UpdatedAt: core.Now(),
		},
	}
	return h.HandlerBase.Insert(c, &doc)
}
`

var findOne = `func (h *Handler) FindOne(c echo.Context) error {
	doc := Model{}
	return h.HandlerBase.FindOne(c, nil, &doc)
}
`

var update = `func (h *Handler) Update(c echo.Context) error {
	doc := Model{
		%s.Model{
			UpdatedAt: core.Now(),
		},
	}
	return h.HandlerBase.UpdateOne(c, &doc)
}
`

var delete = `func (h *Handler) Delete(c echo.Context) error {
	return h.HandlerBase.DeleteOne(c, &Model{})
}
`

func GenerateHandler(packageName string, config Config) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	// has any route? add import if yes
	hasRoute := false
	for _, r := range config.Routes {
		if len(r.Children) > 0 {
			hasRoute = true
			break
		}
	}

	if hasRoute {
		buf = append(buf, "import (")
		buf = append(buf, `	"github.com/labstack/echo/v4"`)
		buf = append(buf, ")\n")
	}

	buf = append(buf, "type IHandler interface {")
	for _, r := range config.Routes {
		for _, c := range r.Children {
			buf = append(buf, fmt.Sprintf(`	%s(c echo.Context) error`, c.Handler))
		}
	}
	buf = append(buf, "}\n")

	buf = append(buf, "///////////////////////////////////////////////////////////////////\n")

	return strings.Join(buf, "\n")
}

func GenerateHandlerService(packageName string) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, "\t\"github.com/labstack/echo/v4\"\n")
	buf = append(buf, fmt.Sprintf("\t\"app/generated/%s\"", packageName))
	buf = append(buf, ")\n")

	buf = append(buf, "type Handler struct {")
	buf = append(buf, fmt.Sprintf("\trest.HandlerBase"))
	buf = append(buf, fmt.Sprintf("\t%s.IHandler", packageName))
	buf = append(buf, "\tRepo Repository")
	buf = append(buf, "}")

	// Generate all, although some handlers will not be used
	// to ignore linting error of imported and not used
	buf = append(buf, fmt.Sprintf(find, strings.Title(packageName)))
	buf = append(buf, fmt.Sprintf(insert, strings.Title(packageName)))
	buf = append(buf, fmt.Sprintf(findOne, strings.Title(packageName)))
	buf = append(buf, fmt.Sprintf(update, strings.Title(packageName)))
	buf = append(buf, fmt.Sprintf(delete, strings.Title(packageName)))

	return strings.Join(buf, "\n")
}
