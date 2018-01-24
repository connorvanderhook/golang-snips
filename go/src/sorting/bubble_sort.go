package main

// BubbleSort implements bubble sort algo
func BubbleSort(items []int) []int {
	for i := len(items) - 1; i >= 0; i-- {
		for j := 1; j <= i; j++ {
			if items[j-1] > items[j] {
				// Don't need temp in golang
				// temp := items[j-1]
				items[j-1], items[j] = items[j], items[j-1]
				// items[j] = temp
			}
		}
	}
	return items
}
