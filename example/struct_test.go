package example

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os/exec"
	"testing"

	"github.com/Just-maple/structgraph"
	"github.com/Just-maple/structgraph/example/internal/app"
	"github.com/Just-maple/structgraph/example/internal/dao/dao_impls"
	"github.com/Just-maple/structgraph/example/internal/database"
	"github.com/Just-maple/structgraph/example/internal/sdk"
	"github.com/Just-maple/structgraph/example/internal/svc/svc_impls"
)

func MakeApplication() *app.Application {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	cache := &database.Redis{}
	db := &database.SqlStore{DB: &sql.DB{}}
	pay := &svc_impls.Pay{Client: sdk.NewPayClient()}
	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {})
	a := &app.Application{
		Router: handler,
		Server: listener,
		Service: app.Service{
			User: &svc_impls.User{
				UserDao: dao_impls.NewUserDao(),
				DB:      db,
				PaySvc:  pay,
			},
			Book: &svc_impls.Book{
				BookDao: dao_impls.NewBookDao(),
				DB:      db,
				Cache:   cache,
			},
			Pay: pay,
		},
	}
	return a
}

func TestDraw(t *testing.T) {
	a := MakeApplication()
	ret := structgraph.Draw(a)
	_ = ioutil.WriteFile("test.puml", []byte("@startuml\n"+ret+"@enduml"), 0775)
	err := ioutil.WriteFile("test.dot", []byte(ret), 0775)
	if err != nil {
		t.Fatal(err)
	}
	genPng()
}

func genPng() {
	cmd := exec.Command(`/bin/sh`, `-c`, "dot test.dot -T png -o test.png")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", out.String())
}
