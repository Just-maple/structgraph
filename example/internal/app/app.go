package app

import (
	"net"

	"github.com/Just-maple/structgraph/example/internal/svc"
)

type Application struct {
	Server  net.Listener
	Service Service
}

type Service struct {
	User svc.IUser
	Book svc.IBook
	Pay  svc.IPay
}
