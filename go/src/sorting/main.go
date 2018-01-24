package main

import "fmt"

func main() {
	// Bubbler
	items := []int{1, 3, 2, 99, 5, 7, 5, 4, 2}
	fmt.Printf("Before Bubble Sort: %v\n", items)
	items = BubbleSort(items)
	fmt.Printf("After Sort: %v\n", items)

	// Bucketer
	items = []int{1, 3, 2, 99, 5, 7, 5, 4, 2}
	fmt.Printf("Before Bucket Sort: %v\n", items)
	// items = BucketSort(items, len(items))
	fmt.Printf("After Sort: %v\n", items)
}
