package core

const AUTH = "NUUzRTU5NzVCOTQ2NkUwMUQ4Mjg5N0JFMzZGNUIxOTE4NkU5RUFENDU3RUYwNDdERkM1MTc5NDY3MUIyMTI5QQ"

type Params struct {
	/// General paging
	Ids   []string // id=1&id=2 -> [1, 2]
	Page  int
	Size  int
	Sort  string
	Order string

	Query string

	CategoryId string
	ProductId  string
	CreatedBy  string
	AccountId  string

	OnSale bool
}
