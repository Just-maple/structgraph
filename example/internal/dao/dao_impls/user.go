package dao_impls

import (
	dao2 `github.com/Just-maple/structgraph/example/internal/dao`
)

type userDao struct {
}

func (u userDao) GetUserByID() {
	return
}

func NewUserDao() dao2.User {
	return userDao{}
}

var _ dao2.User = &userDao{}
