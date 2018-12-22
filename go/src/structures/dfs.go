package structures

func (g *Graph) DFS(v int) {
	marked := make(map[int]struct{})
	adj := make(map[int]int)
	for a := range g.adjacencies[v] {
		if !marked[a] {
			g.DFS(a)
			g.adjacencies[a] = v
		}
	}
}
