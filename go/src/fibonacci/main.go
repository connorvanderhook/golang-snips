package main

import (
	"flag"
	"fmt"
)

var steps int

func init() {
	flag.IntVar(&steps, "steps", 10, "how many steps in the fibonacci sequence to display")
	flag.Parse()
}

func main() {
	f := fibonacci()
	for i := 0; i < steps; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	x := 1
	y := 0
	return func() int {
		x, y = y, x+y
		return x
	}
}
