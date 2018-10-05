package graph

import (
	"errors"
	"fmt"
	"io"
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

// String implements the Stringer interface // pi - u03c0
// returns v[c: color, D: distance, pi: predecessor, adj: len(adj)]
func (v *Vertex) String() string {
	var p Label = "NIL"
	if v.Predecessor != nil {
		p = v.Predecessor.Label
	}
	return fmt.Sprintf(
		"%s[c: %d, D: %d, π: %s, adj: %d]",
		v.Label, v.color, v.Distance, p, len(v.Adj),
	)
}

// Edge E is a pair of vertices (u, v) ∈ E in Graph G = (V, E)
type Edge struct {
	u, v *Vertex
	w    int
}

// String implements the Stringer interface
// to return (u, v):w(#) - # for weight
func (e *Edge) String() string {
	return fmt.Sprintf(
		"(%s, %s): w(%d)",
		e.u.Label, e.v.Label, e.w,
	)
}

// NewEdge returns a reference to an edge (u, v) ∈ V // u2208
// eg. weight w can be added for cases involving weighted graphs
func NewEdge(u, v *Vertex) *Edge {
	return &Edge{
		u: u,
		v: v,
	}
}

// NewWeightedEdge returns a new edge with a given weight
func NewWeightedEdge(u, v *Vertex, w int) *Edge {
	return &Edge{u, v, w}
}

// Graph G = (V, E)
type Graph struct {
	V map[Label]*Vertex
	E map[*Edge]bool

	VNum, ENum int
}

// NewGraph returns a references to a new Graph G
func NewGraph() *Graph {
	return &Graph{
		V: make(map[Label]*Vertex),
		E: make(map[*Edge]bool),
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
			G.VNum++
			G.V[lu] = u
		}
		u := G.V[lu]
		u.Adj = append(u.Adj, Label(lv))
		if G.V[lv] == nil {
			v := NewVertex(lv)
			G.VNum++
			G.V[lv] = v
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
func BFS(G *Graph, l Label) {
	s := G.V[l]
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
}

// PrintPath returns a slice of vertices identified by theier labels
// on a shortest path from s to v,
// assuming that BFS has already been computed
func PrintPath(out io.Writer, G *Graph, a, b Label) {
	s, v := G.V[a], G.V[b]
	if v == s {
		fmt.Fprintln(out, s)
	} else if v.Predecessor == nil {
		fmt.Fprintf(out, "no path from %s to %v\n", a, b)
	} else {
		PrintPath(out, G, s.Label, v.Predecessor.Label)
		fmt.Fprintln(out, v)
	}
}
