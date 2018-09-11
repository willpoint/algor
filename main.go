package main

import (
	"fmt"

	"github.com/willpoint/algor/pkg/ds"
)

func main() {
	bst := ds.NewBST()
	// this benchmark is to test the impact of the improvement to
	// the BST_Insert() function for n insertions of identical keys
	nn := []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	// nn := []int{4, 9, 11, 3, 12, 1, 8, 32, 91, 44, 31, 13, 1, 15, 30}
	for _, v := range nn {
		bst.Insert(v)
		// bst.Insert(v)
	}
	n := bst.Height()
	fmt.Println(n)
}
