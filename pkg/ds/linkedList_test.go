package ds

import (
	"testing"
)

func TestLinkedList(t *testing.T) {

	ss := []string{"go", "c/c++", "haskell", "javascript", "python"}

	l := NewLinkedList()

	for _, j := range ss {
		l.AddHead(j)
	}

	count := len(ss)
	if count != l.Length {
		t.Errorf("Expected %d, got %d", count, l.Length)
	}

	tv := l.Has("c/c++")
	if tv != true {
		t.Errorf("(*linkedList).Has('c/c++') to return true, got false")
	}
}
