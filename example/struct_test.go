package example

import (
	"io/ioutil"
	"testing"

	"github.com/Just-maple/structgraph"
)

type Application struct {
	service  IService
	service4 IService
	service2 *Service2
	service3 Service2
}
type Service struct {
	sql   IMysql
	redis IRedis
}
type Service2 struct {
	sql   IMysql
	redis IRedis
	wxSdk WxSdk
}
type Mysql struct {
	cn Conn
}

type WxSdk interface {
}

type IService interface{}
type Conn struct{}
type IRedis interface{}
type Redis struct{}
type IMysql interface{}

func MakeApplication() Application {
	cn := Conn{}
	d := &Service{
		sql:   Mysql{cn: cn},
		redis: Redis{},
	}
	s2 := &Service2{
		sql:   Mysql{cn: cn},
		redis: Redis{},
		wxSdk: d,
	}
	a := Application{
		service:  d,
		service2: s2,
		service3: Service2{
			redis: s2,
		},
		service4: Mysql{cn: cn},
	}
	d.redis = a
	return a
}

func TestDraw(t *testing.T) {
	app := MakeApplication()
	ret := structgraph.Draw(app)
	_ = ioutil.WriteFile("test.puml", []byte("@startuml\n"+ret+"@enduml"), 0775)
	_ = ioutil.WriteFile("test.dot", []byte(ret), 0775)
}
