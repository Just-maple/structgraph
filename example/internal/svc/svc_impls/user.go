package svc_impls

import (
	"github.com/Just-maple/structgraph/example/internal/dao"
	"github.com/Just-maple/structgraph/example/internal/database"
	svc2 `github.com/Just-maple/structgraph/example/internal/svc`
)

type User struct {
	UserDao dao.User
	DB      database.Store
}

func (u *User) GetUser() {
	panic("implement me")
}

var _ svc2.IUser = &User{}
