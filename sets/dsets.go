package sets

// Member describes a member of a set
type Member struct {
	Name  string
	Next  *Member
	Group *Set
}

// Set describes the structure of a set object
type Set struct {
	Head, Tail *Member
	List       *Member
}

// MakeSet create a new set with a single member and returns
// a pointer to that set
func MakeSet(x string) *Set {
	s := &Set{}
	m := &Member{x, nil, s}
	s.Head = m
	s.Tail = m
	s.List = m
	return s
}

// Union adds all the members of set u to s by
// setting attributes of u to point to s
func Union(s, u *Set) {
	for m := s.List; m != nil; m = m.Next {
		m.Group = s
		s.Tail.Next = m
		s.Tail = m
	}
	u.Head = s.Head
	u.Tail = nil
	u.List = nil
}
