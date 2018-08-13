package ds

// LinkedList is a doubly linked list
type LinkedList struct {
	E    string
	Prev *LinkedList
	Next *LinkedList
}

// Insert into a linkedlist
func (ll *LinkedList) Insert(e string) {
	nl := &LinkedList{e, ll, nil}
	ll.Next = nl
}

// Delete from a linkedlist
func (ll *LinkedList) Delete() {
	if ll.Prev == nil {
		return
	}
	ll.Prev = ll.Next
}

// Has traverses the list and checks if the given string
// exists
func (ll *LinkedList) Has(e string) bool {
	cur := ll
	for cur != nil {
		if ll.E == e {
			return true
		}
		cur = ll.Next
	}
	return false
}
