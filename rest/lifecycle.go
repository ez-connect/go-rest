package rest

import (
	"github.com/ez-connect/go-rest/db"
)

type LifeCycle struct {
	BeforeFind      func(filter *interface{}, option *db.FindOption, projection interface{}) error
	AfterFind       func(docs interface{}) error
	BeforeFindOne   func(filter, projection interface{}) error
	AfterFindOne    func(doc interface{}) error
	BeforeInsert    func(doc interface{}) error
	AfterInsert     func(doc interface{}) error
	BeforeUpdateOne func(filter, doc interface{}) error
	AfterUpdateOne  func(doc interface{}) error
	BeforeDeleteOne func(filter interface{}) error
	AfterDeleteOne  func(doc interface{}) error
}
