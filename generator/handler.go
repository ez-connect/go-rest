package generator

import (
	"fmt"
	"strings"
)

var find = `
func (h *Handler) Find(c echo.Context) error {
	filter := util.GetFilter(c)
	option := util.GetFindOption(c)
	docs := []Model{}
	return h.Find(c, filter, option, ProjectionMeta, &docs)
}
`

var insert = `
func (h *Handler) Insert(c echo.Context) error {
	doc := Model{}
	return h.Insert(c, &doc)
}
`

var findOne = `
func (h *Handler) FindOne(c echo.Context) error {
	filter, err := util.GetFindOneByIdFilter(c)
	if err != nil {
		return util.HTTPErrorBadRequest(err)
	}

	doc := Model{}
	return h.FindOne(c, filter, nil, &doc)
}

`

var update = `
func (h *Handler) Update(c echo.Context) error {
	filter, err := util.GetFindOneByIdFilter(c)
	if err != nil {
		return util.HTTPErrorBadRequest(err)
	}

	doc := Model{}
	return h.UpdateOne(c, filter, &doc)
}

`

var delete = `
func (h *Handler) Delete(c echo.Context) error {
	filter, err := util.GetFindOneByIdFilter(c)
	if err != nil {
		return util.HTTPErrorBadRequest(err)
	}

	return h.DeleteOne(c, filter)
}

`

func GenerateHandler(packageName string) string {
	buf := []string{}
	buf = append(buf, "\n")
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, "\t\"github.com/ez-connect/go-rest/rest\"")
	buf = append(buf, "\t\"github.com/labstack/echo/v4\"")
	buf = append(buf, ")\n")

	buf = append(buf, "type Handler struct {")
	buf = append(buf, "\trest.HandlerBase")
	buf = append(buf, "\tRepo Repository")
	buf = append(buf, "}\n")

	buf = append(buf, "///////////////////////////////////////////////////////////////////\n")
	buf = append(buf, find)
	buf = append(buf, insert)
	buf = append(buf, findOne)
	buf = append(buf, update)
	buf = append(buf, delete)

	return strings.Join(buf, "\n")
}
