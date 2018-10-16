package graph

import (
	"fmt"

	"github.com/willpoint/algor/heap"
)

// EdgeList is a list of edges
type EdgeList []Edge

// Kruskal procedure produces the minimum spanning tree
// for the Graph G = (V, E).
// it works on weighted graphs
// As with a many other algorithms in this package
// it assumes that a dfs has been performed on the graph
// mst-kruskal(G, w)
// a = 0
// for each vertex v ∈ G.V
//  make-set(v)
// sort the edges of G.E into nondecreasing order by weight w
// for each edge (u, v) ∈ G.E, taken in nondecreasing order by weight
//  if find-set(u) != find-set(v)
//   a = a u {(u, v)}
//   union(u, v)
// return a
func Kruskal(G *Graph) {
	_ = make(map[*Vertex][]Edge)
	_ = make(map[*Vertex][]Edge, G.VNum)
	var edges EdgeList
	for e := range G.E {
		edges = append(edges, e)
	}
	h := heap.NewBinaryHeap(&edges)
	h.Sort() // edges is now sorted in non decreasing order
	for _, e := range edges {
		fmt.Println(e.W)
	}
}

// Get return the weight of the edge
func (e EdgeList) Get(i int) interface{} {
	return e[i].W
}

// Set updates the weight of an edge - not needed
func (e EdgeList) Set(i int, x interface{}) {}

// Smaller returns true if weight at index i is smaller than key
func (e EdgeList) Smaller(i int, key interface{}) bool {
	return key.(int) < e.Get(i).(int)
}

// Swap swaps
func (e *EdgeList) Swap(i, j int) {
	ee := *e
	ee[i], ee[j] = ee[j], ee[i]
}

// Push ...
func (e *EdgeList) Push(x interface{}) {}

// Pop ...
func (e *EdgeList) Pop() interface{} {
	ee := *e
	o, j := ee[0], ee[1:]
	*e = j
	return o
}

// Less ...
func (e EdgeList) Less(i, j int) bool {
	return e[i].W < e[j].W
}

// Len ...
func (e EdgeList) Len() int {
	return len(e)
}
