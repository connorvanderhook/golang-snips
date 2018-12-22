package algorithms

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	nodes = map[int][]int{
		1:  []int{2, 7, 8},
		2:  []int{1, 3, 6},
		3:  []int{2, 4, 5},
		4:  []int{3},
		5:  []int{3},
		6:  []int{2},
		7:  []int{1},
		8:  []int{1, 9, 12},
		9:  []int{8, 10, 11},
		10: []int{9},
		11: []int{9},
		12: []int{8},
	}

	dfsTraversalSolution = []int{1, 2, 7, 8, 3, 6, 9, 12, 4, 5, 10, 11}
)

func Test_SimpleGraphDFS(t *testing.T) {
	graph := &SimpleGraph{Nodes: nodes}
	visited := []int{}

	graph.SimpleGraphDFS(4, map[int]bool{}, func(node int) {
		visited = append(visited, node)
	})

	if len(visited) < len(graph.Nodes) {
		t.Fatalf("Graph DFS search did not search all nodes expected:%d", len(graph.Nodes))
	}
}

func Test_SimpleGraphDFSIterative(t *testing.T) {
	graph := &SimpleGraph{Nodes: nodes}
	visited := []int{}

	graph.SimpleGraphDFSIterative(1, map[int]bool{}, func(node int) {
		visited = append(visited, node)
	})
	fmt.Println(visited)
	if len(visited) < len(graph.Nodes) {
		t.Fatalf("Graph DFS search did not search all nodes expected:%d searched:%d\n", len(graph.Nodes), len(visited))
	}
	if !reflect.DeepEqual(visited, dfsTraversalSolution) {
		t.Fatalf("Graph DFS Search did not follow the correct traversal expected: %v \ngot:%v\n", dfsTraversalSolution, visited)
	}
}
