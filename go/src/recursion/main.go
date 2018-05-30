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
	if n == 1 {
		return n
	}
	fmt.Printf("%d+%d\n", n, recursive(n-1))
	return n + recursive(n-1)
}

func main() {
	fmt.Printf("Answer: %d\n", recursive(i))
}
