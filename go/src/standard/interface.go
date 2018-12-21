package main

import "fmt"

type Interface interface {
	Message()
	Counter()
}

type Poop struct {
	S string
	I int
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (poo Poop) Message() {
	fmt.Println(poo.S)
}

// If you only need to print the count of poop
// No need to make this a pointer func
func (poo Poop) Counter() {
	fmt.Println(poo.I)
}

func main() {
	list := []string{"Hello", "my", "name", "is", "Connor"}
	for x, n := range list {
		i := Poop{n, x}
		i.Message()
		i.Counter()
	}
}
