/*
Package sets provides data structures and algorithms
for sets [and disjoint sets]
*/
package sets

type list struct {
	label string
	next  *list
	rep   *Set
}

// Set is a set of object with a linkedlist
// The attributes head points to the first element of the list
// and tail points to the last element of the list.
// Each element in the list is a member of a the same Set,
// it points to the `next` element,
// and also points back [`prev`] to the Set object.
// The representative of Set is the first element in the list
type Set struct {
	head, tail *list
	List       *list
}

// add adds a new set member to the list
func (s *Set) add(m *list) {
	m.rep = s
	s.tail.next = m
	s.tail = m
}

// Union merges to Sets to form a new Set
func (s *Set) Union(u *Set) {
	c := u.List
	s.add(c)
	for c.next != nil {
		s.add(c.next)
		c = c.next
	}
}

// MakeSet creates a new Set
func MakeSet(e string) *Set {
	ns := &Set{}
	m := &list{e, nil, ns}
	ns.head = m
	ns.tail = m
	ns.List = m
	return ns
}

// DSets is a disjoint set represented as a
// map of Set to struct{}
type DSets map[*Set]*list
