/*
@Time : 2021/6/2 10:46
@Author : sunxy
@File : dijkstra_test.go
@description:
*/
package dijkstra

import (
	"testing"
)

func TestGraph_Dijkstra(t *testing.T) {
	a := &Node{Name: "a"}
	b := &Node{Name: "b"}
	c := &Node{Name: "c"}
	d := &Node{Name: "d"}
	e := &Node{Name: "e"}
	f := &Node{Name: "f"}
	g := &Node{Name: "g"}

	graph := Graph{}
	graph.AddEdge(a, c, 2)
	graph.AddEdge(a, b, 5)
	graph.AddEdge(c, b, 1)
	graph.AddEdge(c, d, 9)
	graph.AddEdge(b, d, 4)
	graph.AddEdge(d, e, 2)
	graph.AddEdge(d, g, 30)
	graph.AddEdge(d, f, 10)
	graph.AddEdge(f, g, 1)

	t.Logf(graph.Dijkstra(a))
}
