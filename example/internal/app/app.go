package app

import (
	"net"
	"net/http"

	"github.com/Just-maple/structgraph/example/internal/svc"
)

type Application struct {
	Router  http.Handler
	Server  net.Listener
	Service Service
}

type Service struct {
	User svc.IUser
	Book svc.IBook
	Pay  svc.IPay
}
