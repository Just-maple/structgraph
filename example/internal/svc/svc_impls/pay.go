package svc_impls

import (
	"github.com/Just-maple/structgraph/example/internal/sdk"
	svc2 `github.com/Just-maple/structgraph/example/internal/svc`
)

type Pay struct {
	Client sdk.Pay
}

func (u *Pay) Pay() {
	panic("implement me")
}

var _ svc2.IPay = &Pay{}
