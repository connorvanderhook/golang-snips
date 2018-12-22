package algorithms

type Bag struct {
	vals []int
	index int
}

func NewBag() *Bag {
	return &Bag{vals: make([]int, 0)}
}

func (b *Bag) Add(i int) {
	b.vals = append(b.vals, i)
	b.index++
}

func (b *Bag) IsEmpty(i int) bool {
	return b.index == 0
}

func (b *Bag) Size() int {
	return b.index
}

func (b *Bag) Iterate() <-chan int {
	ch := make(chan int)
	go func() {
		for _, val := range b.vals {
			ch <- val
		}
		close(ch)
	}()
	return ch
}
