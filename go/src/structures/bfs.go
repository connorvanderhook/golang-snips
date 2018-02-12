package structures

import "fmt"

func (g *Graph) BFS(node int, f func(*Vertex)) {
	queue := NewQueue()
	v := g.vertices[node]
	queue.enqueue(*v)
	visited := make(map[*Vertex]bool) // or int 0,1
	for {
		if queue.isEmpty() {
			break
		}
		d := queue.dequeue()
		visited[d] = true

		for _, r := range g.adjacencies[d] {
			j := r
			if !visited[j] {
				queue.enqueue(*r)
				visited[j] = true
			}
		}
		if f != nil {
			f(d)
		}
	}
}

func (g *Graph) PrintGraph() {
	fmt.Println("--- Print Graph ---")
	if len(g.vertices) == 0 {
		return
	}

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