package rest

import (
	"github.com/ez-connect/go-rest/db"
	"github.com/ez-connect/go-rest/rest/filter"
)

type LifeCycle struct {
	BeforeFind      func(params filter.Params, filter *interface{}, option *db.FindOption, projection interface{}) error
	AfterFind       func(params filter.Params, docs interface{}) error
	BeforeFindOne   func(params filter.Params, filter, projection interface{}) error
	AfterFindOne    func(params filter.Params, doc interface{}) error
	BeforeInsert    func(params filter.Params, doc interface{}) error
	AfterInsert     func(params filter.Params, doc interface{}) error
	BeforeUpdateOne func(params filter.Params, filter, doc interface{}) error
	AfterUpdateOne  func(params filter.Params, doc interface{}) error
	BeforeDeleteOne func(params filter.Params, filter interface{}) error
	AfterDeleteOne  func(params filter.Params, doc interface{}) error
}
