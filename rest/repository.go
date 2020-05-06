package rest

import "github.com/ez-conne/golang-rest/db"

type RepositoryBase struct {
	Driver db.DatabaseBase
}

type RepositoryInterface interface {
	EnsureIndexs()
}

func (r *RepositoryBase) Init(driver db.DatabaseBase) {
	r.Driver = driver
}
