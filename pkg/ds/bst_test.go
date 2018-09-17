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
		t.Errorf("BST_Max() expects %d, got %d", want, got)
	}
}

func TestBST_InOrderWalk(t *testing.T) {
	type fields struct {
		Root *BNode
		Num  int
	}
	bst := NewBST()
	nn := []int{11, 4, 7, 15, 9, 3, 6}
	for _, v := range nn {
		bst.Insert(v)
	}
	min, max := 3, 15
	if got := bst.InOrderWalk(); got[0] != min && got[len(nn)-1] != max {
		t.Errorf("BST_InOrderWalk() expects min to be %d and max %d, but got %d and %d",
			min, max, got[0], got[len(nn)-1])
	}
}

func BenchmarkBST_OptimizedInsert(b *testing.B) {
	// This benchmark checks the performance of optimized insert
	// by checking the impact on a BST Search by inserting into an
	// initial BST a series of equal numbers and search for the a different
	// number inserted at a later time during the insertion
	bst := NewBST()
	nn := []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 9}
	for _, v := range nn {
		bst.InsertOptimized(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bst.Search(9)
	}
}

func BenchmarkBST_Insert(b *testing.B) {
	// This benchmark checks the performance of optimized insert
	// by checking the impact on a BST Search by inserting into an
	// initial BST a series of equal numbers and search for the a different
	// number inserted at a later time during the insertion
	bst := NewBST()
	nn := []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 9}
	for _, v := range nn {
		bst.Insert(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bst.Search(9)
	}
}

func TestBST_Search(t *testing.T) {
	type fields struct {
		Root *BNode
		Num  int
	}
	bst := NewBST()
	nn := []int{11, 4, 7, 15, 9, 3, 6}
	for _, v := range nn {
		bst.Insert(v)
	}
	if got := bst.Search(nn[2]); got.Key != nn[2] {
		t.Errorf("Search() expected %d, got %d", nn[2], got.Key)
	}
}

func TestBST_Delete(t *testing.T) {
	type fields struct {
		Root *BNode
		Num  int
	}
	bst := NewBST()
	nn := []int{11, 4, 7, 15, 9, 3, 6}
	for _, v := range nn {
		bst.Insert(v)
	}
	before := bst.Search(11)
	err := bst.Delete(11)
	if err != nil {
		t.Errorf("error occured while deleting: %v", err)
	}
	after := bst.Search(11)
	if before == after {
		t.Errorf("Delete() expects before to not equal after")
	}

}
