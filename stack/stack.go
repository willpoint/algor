/*
Package stack implements a dynamic set in which the elements
removed from the set by the delete operation is prespecified.
The element deleted from the set is the one most recently
inserted - `last-in, first-out` or LIFO, policy.
The operations on a stack are INSERT(PUSH), DELETE(POP).
*/
package stack

// Stack ...
type Stack struct {
	top  int
	elem []interface{}
}

// New initializes a stack
func New() *Stack {
	return &Stack{}
}

// Empty returns true if stack is empty, false otherwise
func (s *Stack) Empty() bool {
	return s.top == 0
}

// Push inserts an item into the stack
func (s *Stack) Push(e interface{}) {
	s.top++
	s.elem = append(s.elem, e)
}

// Pop pops the last inserted item from the stack
// returning an false if the stack is empty
func (s *Stack) Pop() (interface{}, bool) {
	if s.Empty() {
		return nil, false
	}
	s.top--
	e := s.elem[s.top]
	s.elem = s.elem[0 : len(s.elem)-1]
	return e, true
}

// Len returns the length of the stack
func (s *Stack) Len() int {
	return s.top
}
