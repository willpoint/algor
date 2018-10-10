package main

import (
	"github.com/willpoint/algor/graph"
)

func main() {

	l := [][2]string{
		[2]string{"b", "d"},
		[2]string{"e", "g"},
		[2]string{"a", "c"},
		[2]string{"h", "i"},
		[2]string{"a", "b"},
		[2]string{"e", "f"},
		[2]string{"b", "c"},
	}

	G := graph.BuildGraph(l)
	V := G.V
	E := G.E
	_, _ = E, V

}
