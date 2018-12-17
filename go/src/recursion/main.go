package main

import (
	"flag"
	"fmt"
)

var i int

func init() {
	flag.IntVar(&i, "i", 10, "get recursize answer for number")
	flag.Parse()
}

func recursive(n int) int {
	fmt.Printf("%d\n", n)

	if n == 1 {
		return n
	}
	return n + recursive(n-1)
}

func main() {
	fmt.Printf("Answer: %d\n", recursive(i))
}
