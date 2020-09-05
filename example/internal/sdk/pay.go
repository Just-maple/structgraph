package sdk

import (
	"context"
	"net/http"

	utils2 "github.com/Just-maple/structgraph/example/internal/utils"
)

type Pay interface {
	Pay(ctx context.Context)
}

type payClient struct {
	Host   string
	Client utils2.HttpClient
}

func (p payClient) Pay(ctx context.Context) {
	panic("implement me")
}
func NewPayClient() Pay {
	return payClient{
		Host:   "test",
		Client: utils2.HttpClient{Client: &http.Client{}},
	}
}

var _ Pay = payClient{}
