# Go Struct Graph
> generate struct relation graph of your application

## Mentions

- [graphviz](http://www.graphviz.org)
- [awalterschulze/gographviz](github.com/awalterschulze/gographviz)


## Usage example

```go
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

func Draw() {
	app := MakeApplication()
	ret := structgraph.Draw(app)

}
```

draw ret with graphviz `dot`


![dot](example/test.png)