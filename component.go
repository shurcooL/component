// Package component is a collection of basic HTML components.
package component

import (
	"fmt"

	"github.com/shurcooL/htmlg"
	"golang.org/x/net/html"
)

// Text component.
type Text string

func (t Text) Render() []*html.Node {
	return []*html.Node{htmlg.Text(string(t))}
}

// List of components.
type List []htmlg.Component

func (l List) Render() []*html.Node {
	var nodes []*html.Node
	for _, c := range l {
		nodes = append(nodes, c.Render()...)
	}
	return nodes
}

// Join components and strings into a single component.
// Valid types are string, htmlg.Component. Panics on other input.
func Join(a ...interface{}) List {
	var list List
	for _, v := range a {
		switch v := v.(type) {
		case htmlg.Component:
			list = append(list, v)
		case string:
			list = append(list, Text(v))
		default:
			panic(fmt.Errorf("Join: unsupported type: %T", v))
		}
	}
	return list
}
