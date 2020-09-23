package generator

import (
	"fmt"
	"strings"
)

var find = `func (h *Handler) Find%s(c echo.Context) error {
	f := filter.Find(c)
	o := filter.Option(c)
	docs := []Model{}
	return h.Find(c, f, o, nil, &docs)
}
`

var insert = `func (h *Handler) Insert%s(c echo.Context) error {
	doc := Model{}
	return h.Insert(c, &doc)
}
`

var findOne = `func (h *Handler) FindOne%s(c echo.Context) error {
	f := filter.FindOne(c)
	if f == nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	doc := Model{}
	return h.FindOne(c, f, nil, &doc)
}
`

var update = `func (h *Handler) Update%s(c echo.Context) error {
	f := filter.FindOne(c)
	if f == nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	doc := Model{}
	return h.UpdateOne(c, f, &doc)
}
`

var delete = `func (h *Handler) Delete%s(c echo.Context) error {
	f := filter.FindOne(c)
	if f == nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return h.DeleteOne(c, f)
}
`

func GenerateHandler(packageName string) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, "\t\"net/http\"\n")

	buf = append(buf, "\t\"github.com/ez-connect/go-rest/rest\"")
	buf = append(buf, "\t\"github.com/ez-connect/go-rest/rest/filter\"")
	buf = append(buf, "\t\"github.com/labstack/echo/v4\"")
	buf = append(buf, ")\n")

	buf = append(buf, "type Handler struct {")
	buf = append(buf, "\trest.HandlerBase")
	buf = append(buf, "\tRepo Repository")
	buf = append(buf, "}\n")

	buf = append(buf, "///////////////////////////////////////////////////////////////////\n")
	buf = append(buf, fmt.Sprintf(find, strings.Title(packageName)))
	buf = append(buf, fmt.Sprintf(insert, strings.Title(packageName)))
	buf = append(buf, fmt.Sprintf(findOne, strings.Title(packageName)))
	buf = append(buf, fmt.Sprintf(update, strings.Title(packageName)))
	buf = append(buf, fmt.Sprintf(delete, strings.Title(packageName)))

	return strings.Join(buf, "\n")
}

func GenerateHandlerExt(packageName string) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, fmt.Sprintf("\t\"app/generated/%s\"", packageName))
	buf = append(buf, ")\n")

	buf = append(buf, "type Handler struct {")
	buf = append(buf, fmt.Sprintf("\t%s.HandlerBase", packageName))
	buf = append(buf, fmt.Sprintf("\tRepo %s.Repository", packageName))
	buf = append(buf, "}\n")

	return strings.Join(buf, "\n")
}
