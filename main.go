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

	n := G.VNum
	fmt.Println(n)
	n1 := G.DFS()
	fmt.Println(n1)

	GT := ds.GraphTranspose(G)
	nt := GT.VNum
	fmt.Println(nt)
	nt1 := GT.DFS()
	fmt.Println(nt1)
}
