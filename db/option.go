package db

type Order string

// const (
// 	Asc  Order = "asc"
// 	Desc Order = "desc"
// )

type FindOption struct {
	Sort  string
	Order string

	Skip  int64
	Limit int64
}
