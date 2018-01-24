package main

import "fmt"

// BucketSort sort the items
func BucketSort(items []int, max int) map[int]int {
	if max < 1000 {
		resp := make(map[int]int, max)
		for _, item := range items {
			resp[item]++
		}
		return resp
	}
	fmt.Printf("- What do we tell death?\n- Not Today.\n")
	return map[int]int{}
}
