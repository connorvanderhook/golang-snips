package algorithms

// Graph is a Bag Strategy implementation of a graph
// Adjacencies are represented in an Adjacency List
type Graph struct {
	Adjacencies map[int]*Bag
	Nodes       map[int][]int
}

// SimpleGraph uses a basic adjacency list structure
type SimpleGraph struct {
	Nodes AdjacencyList
}

// SimpleGraphDFS recursive implementation of a DFS search on a SimpleGraph struct
func (g *SimpleGraph) SimpleGraphDFS(node int, marked map[int]bool, fn func(int)) {
	marked[node] = true
	fn(node)
	for _, n := range g.Nodes[node] {
		if _, ok := marked[n]; !ok {
			g.SimpleGraphDFS(n, marked, fn)
		}
	}
}

// SimpleGraphDFSIterative is an iterative approach to performine a DFS on a SimpleGraph struct
func (g *SimpleGraph) SimpleGraphDFSIterative(node int, marked map[int]bool, fn func(int)) {
	stack := StackArray{}
	stack.Push(node)

	for {
		if stack.IsEmpty() {
			return
		}

		top := stack.Pop()
		if marked[top] {
			continue
		}

		marked[top] = true
		fn(top)

		for _, n := range g.Nodes[top] {
			if !marked[n] {
				stack.Push(n)
				continue
			}
		}
	}

}
