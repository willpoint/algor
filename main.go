package main

import (
	"fmt"

	"github.com/willpoint/algor/pkg/ds"
)

func main() {

	bst1 := ds.NewBST()
	bst2 := ds.NewBST()
	// this benchmark is to test the impact of the improvement to
	// the BST_Insert() function for n insertions of identical keys
	nn := []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	// nn := []int{4, 9, 11, 3, 12, 1, 8, 32, 91, 44, 31, 13, 1, 15, 30}
	for _, v := range nn {
		bst1.Insert(v)
		bst2.OptimizedInsert(v)
	}
	h1 := bst1.Height()
	h2 := bst2.Height()
	fmt.Println(h1, h2)

}
