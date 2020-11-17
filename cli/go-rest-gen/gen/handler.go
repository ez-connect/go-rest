package gen

import (
	"fmt"
	"strings"
)

var initHandler = `func (h *Handler) Init(db db.DatabaseBase, collection string, repo *Repository) {
	var r rest.RepositoryInterface = &repo.RepositoryBase
	h.HandlerBase.Init(db, collection, r)
	repo.EnsureIndexs()
	%s
}
`

var find = `func (h *Handler) Find%s(c echo.Context) error {
	f, err := filter.Find(c, &Model{})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	o := filter.Option(c)
	docs := []Model{}
	return h.Find(c, f, o, nil, &docs)
}
`

var insert = `func (h *Handler) Insert%s(c echo.Context) error {
	doc := Model{
		CreatedAt: core.Now(),
		UpdatedAt: core.Now(),
	}
	return h.Insert(c, &doc)
}
`

var findOne = `func (h *Handler) FindOne%s(c echo.Context) error {
	f, err := filter.FindOne(c, &Model{})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	doc := Model{}
	return h.FindOne(c, f, nil, &doc)
}
`

var update = `func (h *Handler) Update%s(c echo.Context) error {
	f, err := filter.FindOne(c, &Model{})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	doc := Model{
		UpdatedAt: core.Now(),
	}
	return h.UpdateOne(c, f, &doc)
}
`

var delete = `func (h *Handler) Delete%s(c echo.Context) error {
	f, err := filter.FindOne(c, &Model{})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return h.DeleteOne(c, f)
}
`

func GenerateHandler(packageName string, config Config) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, "\t\"net/http\"\n")

	buf = append(buf, "\t\"github.com/ez-connect/go-rest/db\"")
	buf = append(buf, "\t\"github.com/ez-connect/go-rest/core\"")
	buf = append(buf, "\t\"github.com/ez-connect/go-rest/rest/filter\"")
	buf = append(buf, "\t\"github.com/ez-connect/go-rest/rest\"")
	buf = append(buf, "\t\"github.com/labstack/echo/v4\"\n")

	// buf = append(buf, fmt.Sprintf("\t\"app/services/%s\"", packageName))
	for _, v := range config.Import.Handler {
		buf = append(buf, fmt.Sprintf("\t\"%s\"", v))
	}

	buf = append(buf, ")\n")

	buf = append(buf, "type Handler struct {")
	// buf = append(buf, fmt.Sprintf("\t%s.Handler", packageName))
	buf = append(buf, "\trest.HandlerBase")
	buf = append(buf, "}\n")

	buf = append(buf, "///////////////////////////////////////////////////////////////////\n")

	if config.LifeCycle != "" {
		buf = append(buf, fmt.Sprintf(initHandler, fmt.Sprintf("r.RegisterLifeCycle(%s)", config.LifeCycle)))
	} else {
		buf = append(buf, fmt.Sprintf(initHandler, ""))
	}

	// Generate all, although some handlers will not be used
	// to ignore linting error of imported and not used
	buf = append(buf, fmt.Sprintf(find, strings.Title(packageName)))
	buf = append(buf, fmt.Sprintf(insert, strings.Title(packageName)))
	buf = append(buf, fmt.Sprintf(findOne, strings.Title(packageName)))
	buf = append(buf, fmt.Sprintf(update, strings.Title(packageName)))
	buf = append(buf, fmt.Sprintf(delete, strings.Title(packageName)))

	return strings.Join(buf, "\n")
}

func GenerateHandlerService(packageName string) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, fmt.Sprintf("\t\"app/generated/%s\"", packageName))
	buf = append(buf, ")\n")

	buf = append(buf, "type Handler struct {")
	buf = append(buf, fmt.Sprintf("\t%s.Handler", packageName))
	buf = append(buf, "\tRepo Repository")
	buf = append(buf, "}")

	return strings.Join(buf, "\n")
}
