package main

import (
	"fmt"

	"github.com/willpoint/algor/pkg/ds"
)

func main() {

	a := []string{"undershorts", "shoes", "pants"}
	b := []string{"pants", "belt", "shoes"}
	c := []string{"belt", "jacket"}
	d := []string{"jacket"}
	e := []string{"tie", "jacket"}
	f := []string{"shirt", "belt", "tie"}
	g := []string{"socks", "shoes"}
	h := []string{"watch"}
	i := []string{"shoes"}
	G, _ := ds.NewGraph(a, b, c, d, e, f, g, h, i)
	ll := ds.NewLinkedList()
	G.TopSort(ll)

	n := ll.Head
	for n != nil {
		fmt.Printf("%s-->", n.E)
		n = n.Next
	}
	fmt.Println()

}
