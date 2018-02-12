package structures

import (
	"fmt"
)


type Graph struct {
	vertices []*Vertex
	adjacencies map[*Vertex][]*Vertex
}

func (g *Graph) AddVertex(v *Vertex) {
	g.vertices = append(g.vertices, v)
}

func (g *Graph) AddEdge(v1, v2 *Vertex) {
	if g.adjacencies == nil {
		g.adjacencies = make(map[*Vertex][]*Vertex)
	}
	g.adjacencies[v1] = append(g.adjacencies[v1], v2)
	g.adjacencies[v2] = append(g.adjacencies[v2], v1)
}

// AddEdge adds an edge to the graph
func (g *Graph) String() {
	s := ""
	for i := 0; i < len(g.vertices); i++ {
		s += g.vertices[i].String() + " -> "
		near := g.adjacencies[g.vertices[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
}

func (g *Graph) FillGraph() {
	nA := Vertex{0, false, []*Vertex{}}
	nB := Vertex{1, false, []*Vertex{}}
	nC := Vertex{2, false, []*Vertex{}}
	nD := Vertex{4, false, []*Vertex{}}
	nE := Vertex{5, false, []*Vertex{}}
	nF := Vertex{7, false, []*Vertex{}}

	g.AddVertex(&nA)
	g.AddVertex(&nB)
	g.AddVertex(&nC)
	g.AddVertex(&nD)
	g.AddVertex(&nE)
	g.AddVertex(&nF)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nE, &nF)
	g.AddEdge(&nD, &nA)
}