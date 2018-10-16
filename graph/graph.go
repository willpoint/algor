package graph

import (
	"errors"
	"fmt"
	"io"

	"github.com/willpoint/algor/list"
)

type color int

const (
	white color = iota // first visit
	gray               // adjacent vertices are still being visited
	black              // all adjacent vertices have been visited
)

var (
	// ErrVertexExists ...
	ErrVertexExists = errors.New("vertex already exists")
)

// Label identifying a vertexs
type Label string

// Vertex is a vertex of graph G = (V, E)
type Vertex struct {
	Adj         []Label
	Label       Label
	Predecessor *Vertex
	Distance    int
	dstamp      int
	fstamp      int
	color       color
}

// NewVertex creates a new Vertex
func NewVertex(l Label) *Vertex {
	return &Vertex{
		Label: l,
		color: white,
	}
}

// String implements the Stringer interface // pi - `u03c0`
func (v *Vertex) String() string {
	var p Label = "NIL"
	if v.Predecessor != nil {
		p = v.Predecessor.Label
	}
	return fmt.Sprintf(
		"%s[c: %d, D: %d, π: %s, adj: %d, dstamp: %d, fstamp: %d]",
		v.Label, v.color, v.Distance, p, len(v.Adj), v.dstamp, v.fstamp,
	)
}

// Edge E is a pair of vertices (u, v) ∈ E in Graph G = (V, E)
type Edge struct {
	U, V *Vertex
	W    int
}

// String implements the Stringer interface
// to return (u, v):w(#) - # for weight
func (e Edge) String() string {
	return fmt.Sprintf(
		"(%s, %s), w(%d)",
		e.U.Label, e.V.Label, e.W,
	)
}

// NewEdge returns a reference to an edge (u, v) ∈ V // `u2208`
// eg. weight w can be added for cases involving weighted graphs
func NewEdge(u, v *Vertex) Edge {
	return Edge{
		U: u,
		V: v,
	}
}

// NewWeightedEdge returns a new edge with a given weight
func NewWeightedEdge(u, v *Vertex, w int) Edge {
	return Edge{
		U: u,
		V: v,
		W: w,
	}
}

// Graph G = (V, E)
type Graph struct {
	V map[Label]*Vertex
	E map[Edge]bool

	VNum, ENum int
}

// NewGraph returns a references to a new Graph G
func NewGraph() *Graph {
	return &Graph{
		V: make(map[Label]*Vertex),
		E: make(map[Edge]bool),
	}
}

// String implements the Stringer interface for Graph G
// with the assumption that a BFS, or DFS is already computed
func (G *Graph) String() string {
	var s string
	s += fmt.Sprintln("---Vertex Set---")
	for _, v := range G.V {
		s += fmt.Sprintln(v)
	}
	s += "\n"
	s += fmt.Sprintln("---Edge Set---")
	for e := range G.E {
		s += fmt.Sprintln(e)
	}
	return s
}

// BuildGraph initializes a  Graph G = (V, E) with
// a slice of slice of strings
// for each slice the first element represents u and the
// second represents v
func BuildGraph(pairs [][2]string) *Graph {
	G := NewGraph()
	for _, p := range pairs {
		lu, lv := Label(p[0]), Label(p[1])
		// if lv(2nd element of the pair) has a zero
		// value then lv points to no other vertex
		// we add to Vertex Set and continue
		if lv == Label("") {
			G.V[lu] = NewVertex(lu)
			G.VNum++
			continue
		}
		if G.V[lu] == nil {
			u := NewVertex(lu)
			G.VNum++
			G.V[lu] = u
		}
		G.V[lu].Adj = append(G.V[lu].Adj, Label(lv))
		if G.V[lv] == nil {
			v := NewVertex(lv)
			G.VNum++
			G.V[lv] = v
			edge := NewEdge(G.V[lu], v)
			G.E[edge] = true
		} else {
			edge := NewEdge(G.V[lu], G.V[lv])
			G.E[edge] = true
		}
		G.ENum++
	}
	return G
}

// BuildWeightedGraph is initializes a  Graph G = (V, E)
// with weighted edges
func BuildWeightedGraph(pairs []struct {
	Pair   [2]string `json:"pair"`
	Weight int       `json:"weight"`
}) *Graph {
	G := NewGraph()
	for _, p := range pairs {
		lu, lv, lw := Label(p.Pair[0]), Label(p.Pair[1]), p.Weight
		if G.V[lu] == nil {
			u := NewVertex(lu)
			G.V[lu] = u
			G.VNum++
		}
		u := G.V[lu]
		u.Adj = append(u.Adj, Label(lv))
		if G.V[lv] == nil {
			v := NewVertex(lv)
			G.V[lv] = v
			G.VNum++
			edge := NewWeightedEdge(u, v, lw)
			G.E[edge] = true
		} else {
			edge := NewWeightedEdge(u, G.V[lv], lw)
			G.E[edge] = true
		}
		G.ENum++
	}
	return G
}

// Transpose of Graph G, Gt is graph G with all its
// edges reversed Transpose of G = (V, E) is graph Gt = (V, Et)
func Transpose(g *Graph) *Graph {
	Gt := NewGraph()
	for _, ov := range g.V {
		if Gt.V[ov.Label] == nil {
			Gt.V[ov.Label] = NewVertex(ov.Label)
			Gt.VNum++
		}
		u := Gt.V[ov.Label]
		for _, j := range ov.Adj {
			u.Adj = append(u.Adj, j)
			if Gt.V[j] == nil {
				v := NewVertex(j)
				Gt.V[j] = v
				Gt.VNum++
				edge := NewEdge(v, u)
				Gt.E[edge] = true
				Gt.ENum++
			} else {
				edge := NewEdge(Gt.V[j], u)
				Gt.E[edge] = true
				Gt.ENum++
			}
		}
	}
	return Gt
}

// BFS assumes the input graph is represented using adjacency lists
// the result should be the same for each source as the order of
// visit is always mantained
func BFS(G *Graph, l Label) bool {
	s, ok := G.V[l]
	if !ok {
		return false
	}
	s.color = gray
	Q := []*Vertex{s}
	for len(Q) > 0 {
		u := Q[0]
		Q = Q[1:]
		for _, v := range u.Adj {
			t := G.V[v]
			if t.color == white {
				t.color = gray
				t.Distance = u.Distance + 1
				t.Predecessor = u
				Q = append(Q, t)
			}
		}
		u.color = black
	}
	return true
}

// PrintPath returns a slice of vertices identified by theier labels
// on a shortest path from s to v,
// assuming that BFS has already been computed
func PrintPath(out io.Writer, G *Graph, a, b Label) bool {
	s, ok1 := G.V[a]
	v, ok2 := G.V[b]
	if !ok1 && !ok2 {
		return false
	}
	if v == s {
		fmt.Fprintln(out, s)
	} else if v.Predecessor == nil {
		fmt.Fprintf(out, "no path from %s to %v\n", a, b)
	} else {
		PrintPath(out, G, s.Label, v.Predecessor.Label)
		fmt.Fprintln(out, v)
	}
	return true
}

// Diameter of a tree T = (V, E) gives the largest of all
// shortest-path distances in the tree
// it assumes a BFS has already been computed for the graph
func (G *Graph) Diameter() int {
	var max int
	for _, v := range G.V {
		if v.Distance > max {
			max = v.Distance
		}
	}
	return max
}

// DFS strategy searches deeper into the graph whenever
// possible. It explores edges out of the most discovered
// vertex v that still has unexplored edges leaving it
// Once all of v's edges have been explored, the search
// `backtracks` to explore edges leaving the vertex from
// which v was discovered
// If any undiscovered vertices remain, then depth-first search
// selects one of them as a new source, and it repeats the same
// for that source.
// When depth-first search discovers a vertex v during a scan of
// the adjacency list of an already discovered vertex u,
// it records this event by setting v's predecessor attribute
// v{Predecessor} to u. The predecessor subgraph produced may
// be composed of several trees because the search may repeat
// from multiple sources.
// Notes to remember when classifying an edge
// when we first explore an edge (u, v), the color
// of the vertex tells us something about the edge
// 1. `white` indicates a tree edge
// 2. `gray` indicates a back edge,
// 3. `black` indicates a forward or cross edge
func DFS(G *Graph) int {
	var time int
	var dfsVisit func(*Vertex)
	dfsVisit = func(u *Vertex) {
		u.color = gray
		time++
		u.dstamp = time
		for _, j := range u.Adj {
			v := G.V[j]
			if v.color == white {
				v.Predecessor = u
				v.Distance = u.Distance + 1
				dfsVisit(v)
			}
		}
		u.color = black
		time++
		u.fstamp = time
	}
	for _, u := range G.V {
		if u.color == white {
			dfsVisit(u)
		}
	}
	return time
}

// VertexWalk receives a second parameter fn(e *Vertex) that
// is executed for every vertex completely visited
func VertexWalk(G *Graph, fn func(e *Vertex)) int {
	var time int
	var dfsVisit func(*Vertex)
	dfsVisit = func(u *Vertex) {
		u.color = gray
		time++
		u.dstamp = time
		for _, j := range u.Adj {
			v := G.V[j]
			if v.color == white {
				v.Predecessor = u
				v.Distance = u.Distance + 1
				dfsVisit(v)
			}
		}
		u.color = black
		time++
		u.fstamp = time
		fn(u)
	}
	for _, u := range G.V {
		if u.color == white {
			dfsVisit(u)
		}
	}
	return time
}

// EdgeWalk receives a second parameter fn(e *Edge) that
// is executed for every edge during a depth first
// encountered search of a graph G
func EdgeWalk(G *Graph, fn func(e Edge)) int {
	var time int
	var dfsVisit func(*Vertex)
	dfsVisit = func(u *Vertex) {
		u.color = gray
		time++
		u.dstamp = time
		for _, j := range u.Adj {
			v := G.V[j]
			if v.color == white {
				fn(NewEdge(u, v))
				v.Predecessor = u
				v.Distance = u.Distance + 1
				dfsVisit(v)
			} else {
				fn(NewEdge(u, v))
			}
		}
		u.color = black
		time++
		u.fstamp = time
	}
	for _, u := range G.V {
		if u.color == white {
			dfsVisit(u)
		}
	}
	return time
}

// TopoSort of a DAG produces a linear ordering of all vertices
// such that if G contains an edge (u, v), then u appears before v
// in the ordering - maintaining precedence
// A graph with a cycle cannot produce such an ordering
func TopoSort(G *Graph) *list.LinkedList {
	l := list.NewLinkedList()
	VertexWalk(G, func(v *Vertex) {
		l.AddHead(string(v.Label))
	})
	return l
}

// DFStranspose indicates a forward or cross edge
func DFStranspose(G *Graph, t int) int {
	time := t
	var dfsVisit func(*Vertex)
	dfsVisit = func(u *Vertex) {
		u.color = gray
		u.dstamp = time
		time--
		for _, j := range u.Adj {
			v := G.V[j]
			if v.color == white {
				v.Predecessor = u
				v.Distance = u.Distance + 1
				dfsVisit(v)
			}
		}
		u.color = black
		u.fstamp = time
		time--
	}
	for _, u := range G.V {
		if u.color == white {
			dfsVisit(u)
		}
	}
	return time
}

// SCC (G *Graph)
func SCC(G *Graph) *Graph {
	time := DFS(G)
	fmt.Println(G)
	Gt := Transpose(G)
	time = DFStranspose(Gt, time)
	fmt.Println(Gt)
	return Gt
}
