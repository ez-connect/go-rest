package generator

import (
	"fmt"
	"strings"
)

const settings = `
model:
  name: %s
  attributes:
    - name: name
      type: string
routes:
  - path: /%ss
    children:
      - method: GET
        path: ""
        handler: Find%s
      - method: POST
        path: ""
        handler: Insert%s
      - method: GET
        path: "/:id"
        handler: FindOne%s
      - method: PUT
        path: "/:id"
        handler: Update%s
      - method: DELETE
        path: "/:id"
        handler: Delete%s
`

func GenerateSettings(packageName string) string {
	upper := strings.Title(packageName)
	return fmt.Sprintf(settings, packageName, packageName,
		upper, upper, upper, upper, upper,
	)
}
