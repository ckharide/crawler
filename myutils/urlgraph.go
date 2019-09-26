package myutils

import (
	"fmt"
)

type Vertex struct {
	URL string
}

type Graph struct {
	edges    map[Vertex][]*Vertex
	V        int
	E        int
	vertexes []Vertex
}

type Edge struct {
	u, v Node
}

// AddNode : Exporting Add Node.
func (g *Graph) AddNode(v Vertex) {
	g.vertexes = append(g.vertexes, v)
	fmt.Println("Adding Node ", v.URL)
}

// AddEdge : Exporting Add Edge.
func (g *Graph) AddEdge(n1, n2 *Vertex) {
	if g.edges == nil {
		g.edges = make(map[Vertex][]*Vertex)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}

// Build Graph : Building a graph
func (g *Graph) BuildGraph(V int) *Graph {
	fmt.Println("Building Graph strcuture ")
	g.vertexes = make([]Vertex, V, )
	return g
}

func (g *Graph) GetNodes() int {
	return len(g.vertexes)
}

func (g *Graph) Print() {
	for k, v := range g.edges {
		fmt.Printf("key[%s] ", k.URL)
		fmt.Println("Values ===>")
		for _, p := range v {
			fmt.Printf("%+v\n", p)
		}
	}
}
