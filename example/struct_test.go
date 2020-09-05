package example

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net"
	"os/exec"
	"testing"

	"github.com/Just-maple/structgraph"
	"github.com/Just-maple/structgraph/example/internal/app"
	"github.com/Just-maple/structgraph/example/internal/dao/dao_impls"
	"github.com/Just-maple/structgraph/example/internal/database"
	"github.com/Just-maple/structgraph/example/internal/sdk"
	"github.com/Just-maple/structgraph/example/internal/svc/svc_impls"
	_ "github.com/go-sql-driver/mysql"
)

func MakeApplication() app.Application {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	src := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "Aa123456", "0.0.0.0", 3306, "mysql")
	sqlDB, err := sql.Open("mysql", src)
	if err != nil {
		panic(err)
	}
	db := &database.SqlStore{DB: sqlDB}
	pay := &svc_impls.Pay{Client: sdk.NewPayClient()}
	a := app.Application{
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
	fmt.Printf("%s", out.String()) //输出执行结果
}
