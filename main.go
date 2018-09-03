package main

import (
	"fmt"

	"github.com/willpoint/algor/pkg/ds"
)

func main() {

	bst := ds.NewBST()
	nn := []int{11, 4, 7, 15, 9, 3, 6}
	for _, v := range nn {
		// fmt.Println(v)
		bst.Insert(v)
	}
	curr := bst.Root
	for curr != nil {
		fmt.Println(curr.Key)
		curr = curr.Left
	}

}
