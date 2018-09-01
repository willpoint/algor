package ds

import (
	"testing"
)

func TestBST_Insert(t *testing.T) {
	bst := NewBST()
	nn := []int{11, 4, 7, 15, 9, 3, 6}
	have := len(nn)
	for _, v := range nn {
		bst.Insert(v)
	}
	if got := bst.Num; got != have {
		t.Errorf("TestBST_Insert() got %d, expected %d", got, have)
	}
}
