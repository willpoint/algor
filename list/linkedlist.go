package list

// LinkedList ...
type LinkedList struct {
	Head   *LNode
	Length int
}

// NewLinkedList ...
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// LNode is a doubly linked list
type LNode struct {
	E    string
	Prev *LNode
	Next *LNode
}

// NewNode creates a new LNode element
func NewNode(e string) *LNode {
	return &LNode{
		E: e,
	}
}

// AddHead adds a new element to the head of the LNode
// if head is does not exist then the first element becomes
// the element at head
func (ll *LinkedList) AddHead(e string) {
	head := ll.Head
	nn := NewNode(e)
	if head == nil {
		ll.Head = nn
	} else {
		nn.Next = head
		head.Prev = nn
		ll.Head = nn
	}
	ll.Length++
}

// Get node element with e
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

// Has traverses the list and checks
// if the given string exists
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

// String implements a Stringer interface
func (ll *LinkedList) String() (s string) {
	s = "|"
	node := ll.Head
	for node != nil {
		s += node.E + "|"
		node = node.Next
	}
	return s
}
