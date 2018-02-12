package structures

import (
"fmt"
)

// Vertex is a single instance in a graph
type Vertex struct {
	val int
	visited bool
	adj []*Vertex
}

func (v *Vertex) String() string {
	return fmt.Sprintf("%d", v.val)
}

