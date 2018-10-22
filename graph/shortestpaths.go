package graph

import (
	"github.com/willpoint/algor/heap"
)

// initSingleSource initializes the shortest-path
// estimates for all vertices in graph G, and sets
// the estimate for vertex labeled s to 0
// The estimate is an upper bound on the weight of a shortest
// path from source s to v - which would be set to the largest
// int value depending on the target architecture
// after initialization - v.Distance = max for all v ∈ V - {src}
// This is achieved in an order 0(V)-time, where V is the number of Vertices
func (G *Graph) initSingleSource(src Label) {
	for _, v := range G.V {
		v.Distance = int(^uint(0) >> 1)
	}
	G.V[src].Distance = 0
}

// relaxing an edge (u, v) is done based on a condition
// that tells if we can improve the shortest path to v
// found so far by going through u and, if the condition
// passes, update v.Distance and v.Predecessor
// This is achieved in an order 0(1)-time
func (G *Graph) relax(e Edge) {
	w := G.E[e]
	improved := e.V.Distance > e.U.Distance+w
	if improved {
		e.V.Distance = e.U.Distance + w
		e.V.Predecessor = e.U
	}
}

// vertices satisfy heap.Heaper interface for binary priority queue
type vertices []*Vertex

func (v vertices) Len() int                            { return len(v) }
func (v vertices) Smaller(i int, key interface{}) bool { return v[i].Distance < key.(*Vertex).Distance }
func (v vertices) Less(i, j int) bool                  { return v[i].Distance < v[j].Distance }
func (v vertices) Get(i int) interface{}               { return v[i] }

func (v *vertices) Pop() interface{} {
	n := *v
	i, j := n[0], n[1:]
	*v = j
	return i
}

func (v *vertices) Push(x interface{}) {
	n := *v
	n = append(n, x.(*Vertex))
}

func (v *vertices) Swap(i, j int) {
	n := *v
	n[i], n[j] = n[j], n[i]
}

func (v *vertices) Set(i int, x interface{}) {
	n := *v
	n[i] = x.(*Vertex)
}

// Dijkstra solves the shortest-paths problem on a weighted,
// directed graph G, for which all edge weights are nonnegative
// in order of 0((V+E)log V) using a binary min-heap priority queue
// It maintains a map S of vertices whose final shortest-path weights
// from the source src have already been determined
func Dijkstra(G *Graph, src Label) []*Vertex {
	G.initSingleSource(src)
	S := []*Vertex{}
	var V vertices
	for _, v := range G.V {
		V = append(V, v)
	}
	Q := heap.NewBinaryHeap(&V)
	Q.BuildMinHeap()
	for !Q.Empty() {
		min, _ := Q.ExtractMin()
		u := min.(*Vertex)
		S = append(S, u)
		for _, j := range u.Adj {
			v := G.V[j]
			G.relax(Edge{u, v})
		}
		Q.MinHeapify(0)
	}
	return S
}
