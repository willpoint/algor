package ds

// LinkedList ...
type LinkedList struct {
	Head   *LNode
	Length int
}

// NewLinkedList ...
func NewLinkedList() *LinkedList {
	head := &LNode{}
	return &LinkedList{head, 0}
}

// LNode is a doubly linked list
type LNode struct {
	E    string
	Prev *LNode
	Next *LNode
}

// NewLNode creates a new LNode element
func NewLNode(e string) *LNode {
	return &LNode{
		E: e,
	}
}

// AddHead adds a new element to the head of the LNode
func (ll *LinkedList) AddHead(e string) {
	old := ll.Head
	nl := NewLNode(e)
	ll.Head = nl
	ll.Head.Next = old
	ll.Length++
}

// Get ...
func (ll *LinkedList) Get(e string) *LNode {
	cur := ll.Head
	for cur != nil {
		if cur.E == e {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

// Has traverses the list and checks if the given string
// exists
func (ll *LinkedList) Has(e string) bool {
	cur := ll.Head
	for cur != nil {
		if cur.E == e {
			return true
		}
		cur = cur.Next
	}
	return false
}
