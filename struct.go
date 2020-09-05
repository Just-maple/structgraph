package structgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
	"strings"

	"github.com/awalterschulze/gographviz"
)

type (
	Option func(d *drawer)

	drawer struct {
		graph         *gographviz.Graph
		scopes        []string
		depth         int
		edgeStore     map[string]bool
		showItfMethod bool

		fn []func()
	}
)

func GenPuml(app interface{}, opts ...Option) (err error) {
	ret := Draw(app, opts...)
	err = ioutil.WriteFile("test.puml", []byte("@startuml\n"+ret+"@enduml"), 0775)
	return
}

func Draw(in interface{}, opts ...Option) string {
	graphAst, _ := gographviz.ParseString(`digraph G {}`)
	graph := gographviz.NewGraph()
	_ = gographviz.Analyse(graphAst, graph)
	v := reflect.ValueOf(in)
	dw := &drawer{graph: graph, edgeStore: map[string]bool{}}
	b, _ := getModBase()
	if len(b) > 0 {
		dw.scopes = append(dw.scopes, b)
	}
	for _, o := range opts {
		o(dw)
	}
	dw.draw(context.Background(), "", v, 0)
	for _, f := range dw.fn {
		f()
	}
	return graph.String()
}

func (d *drawer) addField(structName string, fieldName string, level int) {
	d.fn = append(d.fn, func() {
		size := d.getSize(level)
		d.addPNode(fmt.Sprintf(`"cluster_%s"`, structName), structName,
			map[string]string{
				"fontsize": fmt.Sprintf("%v", size),
				"margin":   fmt.Sprintf("%v", size*0.01),
				"shape":    "tab",
			})
		d.addPNode(fmt.Sprintf(`"cluster_%s"`, structName), structName+":"+fieldName,
			map[string]string{
				"label":    fieldName,
				"shape":    "box",
				"fontsize": fmt.Sprintf("%v", size*0.85),
				"margin":   fmt.Sprintf("%v", size*0.01),
			})
	})
	d.addEdge(structName, structName+":"+fieldName, map[string]string{"arrowhead": "dot"})
}

func (d *drawer) addImpl(itf reflect.Type, implType string, level int) {
	size := d.getSize(level)
	label := itf.String()
	if d.showItfMethod {
		label += ":"
		for i := 0; i < itf.NumMethod(); i++ {
			m := itf.Method(i)
			label += "\n" + m.Name + " : " + m.Type.String()
		}
	}
	d.addNode(itf.String(), map[string]string{
		"shape":  "component",
		"margin": fmt.Sprintf("%v", size*0.01),
		"label":  fmt.Sprintf(`"%s"`, label)}, level)
	d.addNode(implType, nil, level)
	d.addEdge(itf.String(), implType, map[string]string{"label": "impl", "style": "dashed"})
}

func (d *drawer) addEdge(from string, to string, attrs map[string]string) {
	from = fmt.Sprintf(`"%s"`, from)
	to = fmt.Sprintf(`"%s"`, to)

	if !d.edgeStore[from+"_"+to] {
		d.edgeStore[from+"_"+to] = true
		_ = d.graph.AddEdge(from, to, true, attrs)
	}
}

func (d *drawer) addNode(nodeName string, attrs map[string]string, level int) {
	d.fn = append(d.fn, func() {
		if attrs == nil {
			attrs = map[string]string{}
		}
		attrs["fontsize"] = fmt.Sprintf("%v", d.getSize(level))
		d.addPNode("G", nodeName, attrs)
	})
}

func (d *drawer) getSize(level int) float64 {
	return math.Pow(float64(d.depth-level), 2) + 15
}

func (d *drawer) addPNode(parent string, nodeName string, attrs map[string]string) {
	nodeName = fmt.Sprintf(`"%s"`, nodeName)
	_ = d.graph.AddNode(parent, nodeName, attrs)
}

func (d *drawer) addSubGraph(nodeName, pkg string, level int) {
	d.fn = append(d.fn, func() {
		if len(d.scopes) > 0 {
			pkg = strings.Replace(pkg, d.scopes[0], "", -1)
		}
		size := d.getSize(level)
		_ = d.graph.AddSubGraph("G",
			fmt.Sprintf(`"cluster_%s"`, nodeName),
			map[string]string{
				"label":    fmt.Sprintf(`"%s"`, pkg),
				"fontsize": fmt.Sprintf(`%v`, size),
				"style":    "dashed",
				"margin":   fmt.Sprintf(`%v`, size),
				"labelloc": "t",
			})
	})
}

func (d *drawer) draw(ctx context.Context, parent string, v reflect.Value, level int) {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if len(parent) > 0 {
		d.addEdge(parent, v.Type().String(), nil)
	}

	for v.Kind() == reflect.Interface {
		t := v.Type()
		v = v.Elem()
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		d.addImpl(t, v.Type().String(), level)
	}

	d.addNode(
		v.Type().String(),
		nil,
		level,
	)

	if v.Kind() != reflect.Struct {
		return
	}

	if len(d.scopes) > 0 {
		var inScope bool
		for _, s := range d.scopes {
			if strings.Contains(v.Type().PkgPath(), s) {
				inScope = true
				break
			}
		}
		if !inScope {
			return
		}
	}

	d.addSubGraph(v.Type().String(), v.Type().PkgPath(), level)

	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name

		d.addField(v.Type().String(), fieldName, level)

		if v.Field(i).IsZero() {
			continue
		}

		tmp := v.Field(i)

		if tmp.Kind() == reflect.Ptr {
			tmp = tmp.Elem()
		}

		if tmp.Kind() != reflect.Interface && tmp.Kind() != reflect.Struct {
			continue
		}

		next := v.Type().String() + ":" + fieldName

		if ctx.Value(v.Field(i)) != nil {
			continue
		}
		nctx := context.WithValue(ctx, v.Field(i), struct{}{})

		if d.depth < level+1 {
			d.depth = level
		}
		d.draw(nctx, next, v.Field(i), level+1)
	}
}
