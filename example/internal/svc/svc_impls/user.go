package svc_impls

import (
	"github.com/Just-maple/structgraph/example/internal/dao"
	"github.com/Just-maple/structgraph/example/internal/database"
	"github.com/Just-maple/structgraph/example/internal/svc"
)

type User struct {
	UserDao dao.User
	DB      database.Store
	PaySvc  svc.IPay
}

func (u *User) GetUser() {
	panic("implement me")
}

var _ svc.IUser = &User{}
