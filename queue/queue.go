/*
Package queue implements INSERT(ENQUEUE) and DELETE(DEQUEUE)
operations on the underlying slice. The first element inserted
is the first that is returned during a dequeue op
The queue has a head and a tail, when an element is enqueued
it takes it's position at the tail of the queue. while the element
dequeued is always at the head of the queue
*/
package queue

// Queue struct
type Queue struct {
	head, tail int
	elem       []interface{}
}

// New initializes a Queue
func New() *Queue {
	return &Queue{}
}

// Empty returns true if queue is empty and false otherwise
func (q *Queue) Empty() bool {
	return q.head == q.tail
}

// Enqueue inserts a new element into the queue
func (q *Queue) Enqueue(e interface{}) {
	q.elem = append(q.elem, e)
	q.tail = len(q.elem)
}

// Dequeue removes the element at head from the queue
// returns false if queue is empty
func (q *Queue) Dequeue() (interface{}, bool) {
	if q.Empty() {
		return nil, false
	}
	e := q.elem[q.head]
	q.elem = q.elem[1:]
	q.tail = len(q.elem)
	return e, true
}

// Len returns the length of the queue
func (q *Queue) Len() int {
	return q.tail
}
