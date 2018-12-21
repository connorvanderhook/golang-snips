package main

import (
	"bytes"
	"fmt"
)

func main() {
	// make(type, length, capacity)
	idx1 := make([]int, 0, 10)
	idx1 = append(idx1, 1)
	// prints [1]
	fmt.Println(idx1)

	// make(type, length, capacity)
	idx2 := make([]int, 2, 10)
	idx2 = append(idx2, 1)
	// prints [0,0,1]
	fmt.Println(idx2)

	// make(type, length, capacity)
	idx3 := make([]int, 2)
	idx3 = append(idx3, 1)
	// prints [0,0,1]
	// fmt.Println(idx3)
	splitsDemo()

	appendDemo()

	// append(x, 1) <- Maginc
	// First, it can append to a nil slice, making it spring into existence while appending.
	// Second, if the remaining capacity is not sufficient for appending new values,
	// append() automatically takes care of allocating a new array and copying the old content over.
}

func splitsDemo() {
	fmt.Println("\nSplits demo")
	a := []byte("a, b, c")
	b := bytes.Split(a, []byte(","))
	fmt.Printf("a before changing b[0][0]: %q\n", a)
	b[0][0] = byte('*')
	fmt.Printf("a after changin b[0][0]: %q\n", a)

	fmt.Printf("b[1] before appending to b[0]: %q\n", b[1])
	b[0] = append(b[0], 'd', 'e', 'f')
	fmt.Printf("b[1] after appending to b[0]:  %q\n", b[1])
	fmt.Printf("a after appending to b[0]: %q\n", a)
}

func appendDemo() {
	fmt.Println("\nAppend demo")
	s1 := make([]int, 2, 4)
	s1[0] = 1
	s1[1] = 2
	fmt.Printf("Initial address and value: %p: %[1]v\n", s1)
	// Append numbers to a slice within capacity.
	s1 = append(s1, 3, 4)
	// Note the same address as before.
	fmt.Printf("After first append:        %p: %[1]v\n", s1)

	// Note the changed address. Append allocated a new, larger array.
	s1 = append(s1, 5)
	fmt.Printf("After second append:       %p: %[1]v\n", s1)
}
