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

	count := 0
	node := l.Head
	for node != nil {
		if node.Next != nil {
			count++
		}
		node = node.Next
	}

	expected := len(ss)
	if count != expected {
		t.Errorf("Expected %d, got %d", expected, count)
	}

	tv := l.Has("c/c++")
	if tv != true {
		t.Errorf("(*linkedList).Has('c/c++') to return true, got false")
	}
}