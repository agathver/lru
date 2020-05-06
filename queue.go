package lru

type Queue struct {
	size int
	head *Entry
	tail *Entry
}

func NewQueue() *Queue {
	return &Queue{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Offer(value interface{}) {
	e := NewEntry(value)

	if q.head == nil && q.tail == nil {
		q.head = e
		q.tail = e
	} else {
		e.prev = q.tail
		q.tail.next = e
		q.tail = e
	}
	q.size++
}

func (q *Queue) Poll() interface{} {
	if q.size == 0 {
		return nil
	}

	e := q.head
	q.head = q.head.next

	q.size--
	if q.size == 0 {
		q.tail = nil
	}

	return e.value
}
