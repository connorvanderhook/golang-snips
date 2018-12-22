package algorithms

type Stack interface {
	Push(i interface{})
	Pop() interface{}
	IsEmpty() bool
	Size() int
	Top() int
	// Iterate me maybe?
}

type StackArray struct {
	Items []interface{}
	// Not going to implement this until needed
	// Index int
}

func NewStackArray() StackArray {
	return StackArray{Items: make([]interface{}, 0)}
}

func (s *StackArray) Push(v int) {
	s.Items = append(s.Items, v)
}

func (s *StackArray) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	item, items := s.Items[0], s.Items[1:]
	s.Items = items
	return item.(int)
}

func (s *StackArray) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *StackArray) Size() int {
	return len(s.Items)
}

func (s *StackArray) Top() int {
	return s.Items[len(s.Items)-1].(int)
}
