package svc_impls

import (
	"github.com/Just-maple/structgraph/example/internal/dao"
	"github.com/Just-maple/structgraph/example/internal/database"
	"github.com/Just-maple/structgraph/example/internal/svc"
)

type Book struct {
	DB      database.Store
	Cache   database.Cache
	BookDao dao.Book
}

func (u *Book) GetBook() {
	panic("implement me")
}

var _ svc.IBook = &Book{}
