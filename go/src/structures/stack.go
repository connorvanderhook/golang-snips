package structures

import "log"

// Stack is a stack
type Stack struct {
	Size  int   // size of stack
	First *Node // top of stack
}

// Node is a node
type Node struct {
	Val  int
	Next *Node
}

// NewStack constructs a new stack
func NewStack() *Stack {
	return &Stack{Size: 0, First: nil}
}

//IsEmpty returns if is empty
func (stck *Stack) IsEmpty() bool {
	return stck.First == nil
}

// GetSize is len
func (stck *Stack) GetSize() int {
	return stck.Size
}

// Stack adds an item (int) to stack
func (stck *Stack) Push(x int) {
	tmp := stck.First
	stck.First = &Node{Val: x, Next: tmp}
	stck.Size++
}

func (stck *Stack) Pop() int {
	if stck.IsEmpty() {
		log.Fatal("Stack underflow")
	}
	item := stck.First.Val
	stck.First = stck.First.Next
	stck.Size--
	return item
}

// func main() {
// 	s := NewStack()
// 	println(s.IsEmpty())
// 	println(s.GetSize())
// 	s.Push(1)
// 	s.Push(4)
// 	s.Push(9)
// 	println(s.Pop())
// 	s.Push(7)
// 	println(s.Pop())
// 	println(s.Pop())
// 	println(s.IsEmpty())
// 	println(s.GetSize())
// }
