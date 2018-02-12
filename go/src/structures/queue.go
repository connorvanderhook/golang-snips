package structures

type Queue struct {
	items []*Vertex
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) enqueue(v Vertex) {
	q.items = append(q.items, &v)
}

func (q *Queue) dequeue() *Vertex {
	if q.isEmpty() {
		return nil
	}
	val, vals := q.items[0], q.items[1:]
	q.items = vals
	return val
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) front() *Vertex {
	return q.items[0]
}

func (q *Queue) size() int {
	return len(q.items)
}
