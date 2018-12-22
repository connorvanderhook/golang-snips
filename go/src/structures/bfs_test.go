package structures

import (
	"fmt"
	"testing"
)

func Test_bfs(t *testing.T) {
	g := &Graph{}

	g.FillGraph()

	g.PrintGraph()

	g.BFS(5, func(v *Vertex) {
			fmt.Println("--- Test BFS ---")
			fmt.Printf("%v\n", v)
		}
	)
}