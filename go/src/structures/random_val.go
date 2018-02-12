package structures

import "math/rand"

func randEdges(seed int) []int {
	return []int{rand.Intn(seed)}
}