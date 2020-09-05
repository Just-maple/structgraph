package example

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net"
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
	a := app.Application{
		Server: listener,
		Service: app.Service{
			User: &svc_impls.User{
				UserDao: dao_impls.NewUserDao(),
				DB:      db,
			},
			Book: &svc_impls.Book{
				BookDao: dao_impls.NewBookDao(),
				DB:      db,
			},
			Pay: &svc_impls.Pay{Client: sdk.NewPayClient()},
		},
	}
	return a
}

func TestDraw(t *testing.T) {
	a := MakeApplication()
	ret := structgraph.Draw(a)
	_ = ioutil.WriteFile("test.puml", []byte("@startuml\n"+ret+"@enduml"), 0775)
	_ = ioutil.WriteFile("test.dot", []byte(ret), 0775)
}
