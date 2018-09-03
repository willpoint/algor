package ds

import (
	"testing"
)

func TestBST_Insert(t *testing.T) {
	bst := NewBST()
	nn := []int{11, 4, 7, 15, 9, 3, 6}
	want := 11
	for _, v := range nn {
		bst.Insert(v)
	}

	if got := bst.Root.Key; got != want {
		t.Errorf("TestBST_Insert() got %d, expected %d", got, want)
	}
}

func TestBST_Min(t *testing.T) {
	type fields struct {
		Root *BNode
		Num  int
	}
	bst := NewBST()
	nn := []int{11, 4, 7, 15, 9, 3, 6}
	for _, v := range nn {
		bst.Insert(v)
	}
	want := 3 // min of the array nn
	if got := bst.Min(); got != want {
		t.Errorf("BST_Min() expects %d, got %d", want, got)
	}
}

func TestBST_Max(t *testing.T) {
	type fields struct {
		Root *BNode
		Num  int
	}
	bst := NewBST()
	nn := []int{11, 4, 7, 15, 9, 3, 6}
	for _, v := range nn {
		bst.Insert(v)
	}
	want := 15 // min of the array nn
	if got := bst.Max(); got != want {
		t.Errorf("BST_Min() expects %d, got %d", want, got)
	}
}
