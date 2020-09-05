package dao_impls

import (
	dao2 `github.com/Just-maple/structgraph/example/internal/dao`
)

type bookDao struct {
}

func (u bookDao) GetBookByID() {
	return
}

func NewBookDao() dao2.Book {
	return bookDao{}
}

var _ dao2.Book = &bookDao{}
