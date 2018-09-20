package ds

import (
	"errors"
	"fmt"
)

// Color ...
type Color int

var (
	// ErrVertexExists ...
	ErrVertexExists = errors.New("vertex already exists")
)

const (
	// White represents an undiscovered vertex
	White Color = iota
	// Gray is discovered but becomes a frontier for
	// discovered or undiscovered Vertices
	Gray
	// Black is a discovered vertex whose adjacent vertices
	// if exists have been discovered
	Black
)

// Vertex ...
type Vertex struct {
	Label       string
	Color       Color
	Distance    int
	Predecessor *Vertex

	// Discovered timestamp -
	// for when the vertex is discovered used during DFS
	DStamp int

	// Finished timestamp -
	// for when the edges for the vertex have
	// been completely discovered used for DFS algorithm
	FStamp int

	adj *LinkedList
}

// Edge is a pair of vertex (u, v)
type Edge struct {
	u, v *Vertex
}

// NewEdge returns a reference to an edge (u, v) ∈ V // u2208- unicode
// eg. weight w can be added for cases involving weighted graphs
func NewEdge(u, v *Vertex) *Edge {
	return &Edge{u, v}
}

// Graph G = (V, E) where V is the vertex set
// and E is the edge set containing edges (u, v) ∈ V
type Graph struct {
	V map[string]*Vertex
	E map[*Edge]struct{}

	VNum int // number of vertices in G.V
	ENum int // number of edges in G.E
}

// NewGraph returns a references to a new Graph G
func NewGraph() *Graph {
	return &Graph{
		V: make(map[string]*Vertex),
		E: make(map[*Edge]struct{}),
	}
}

// BuildGraph is a variadic function that initializes a
// Graph G = (V, E) with a slice of slice of strings
// representing the label for each vertex v in G.
// The adjacency list for each vertex u ∈ V is represented
// as elements following string at first index of the slice S[1, 1i)
func BuildGraph(ll ...[]string) (*Graph, error) {

	G := NewGraph()

	for _, j := range ll {
		v := NewVertex(j[0])
		for _, k := range j[1:] {
			v.adj.AddHead(k)
			G.E[NewEdge(v, NewVertex(k))] = struct{}{}
			G.ENum++
		}
		// each adjacency list is composed of
		// j[0] the item itself and j[1:]..., the elements
		// making up the adjlist for j[0]
		G.V[j[0]] = v
		G.VNum++
	}

	return G, nil
}

// V ...
func (v *Vertex) V() *LinkedList {
	return v.adj
}

// VertexCount ...
func (v *Vertex) VertexCount() int {
	return v.adj.Length
}

// NewVertex creates a new Vertex initialized with
// reasonable defaults
func NewVertex(l string) *Vertex {
	return &Vertex{
		Label: l,
		adj:   NewLinkedList(),
		Color: White,
	}
}

// GraphTranspose of Graph G, Gt is graph G with all its
// edges revered
// Transpose of G = (V, E) is graph Gt = (V, Et)
func GraphTranspose(g *Graph) *Graph {
	gt := NewGraph()

	for _, v := range g.V {
		if _, ok := gt.V[v.Label]; !ok {
			gt.V[v.Label] = NewVertex(v.Label)
			gt.VNum++
		}
		curr := v.adj.Head
		for curr != nil {
			_, exists := gt.V[curr.E]
			if !exists {
				gt.V[curr.E] = NewVertex(curr.E)
				gt.VNum++
			}
			u := gt.V[curr.E]
			u.adj.AddHead(v.Label)
			gt.E[NewEdge(u, v)] = struct{}{}
			gt.ENum++
			curr = curr.Next
		}
	}
	return gt
}

// BFS when given a source vertex s, systematically explores the
// edges of G to discover every vertex that is reachable from s
// `It computes the smallest number of edges from s to each
// reachable vertex producing a `breadth-first-tree` with root s
// that contains all reachable vertices`
// It discovers all vertices at distance k from s before discovering
// vertices at distance k + 1
// It uses colors to keep track of the progress ensuring that
// each edge is visited only once, and increments the Distance
// attribute of the vertex v to the min number of edges reachable
// from source s
// It uses queue (Q) to maintain the breadth-first search sequence
func (g *Graph) BFS(label string) error {

	s, present := g.V[label]
	if !present {
		return errors.New("vertex does not exist in graph")
	}

	s.Color = Gray

	Q := []*Vertex{}
	Q = append(Q, s)

	for len(Q) > 0 {

		u := Q[0] // pop the first vertex out of queue Q
		Q = Q[1:] // update queue to account for the popped vertex

		// u.adj is a list of adjacent vertices from u
		node := u.adj.Head
		for node != nil {
			v, present := g.V[node.E]
			if present { // if (u, v) is a member of E
				if v.Color == White {
					v.Color = Gray
					v.Distance = u.Distance + 1
					v.Predecessor = u
					Q = append(Q, v) // add vertex to queue to be visited next
				}
			}
			node = node.Next
		}
		u.Color = Black
	}
	return nil
}

// PrintPath prints out the vertices on a shortest path from s to v
// assuming that BFS has already computed the breadth-first tree
// for src
func (g *Graph) PrintPath(src, dst string) {
	s := g.V[src]
	v := g.V[dst]
	if v == s {
		fmt.Printf("%s(d-%d)\n", s.Label, s.Distance)
	} else if v.Predecessor == nil {
		fmt.Printf("no path from %s to %s exists\n", s.Label, v.Label)
	} else {
		g.PrintPath(s.Label, v.Predecessor.Label)
		fmt.Printf("%s(d-%d)\n", v.Label, v.Distance)
	}
}

// Diameter of a tree T = (V, E) gives the largest of all
// shortest-path distances in the tree
// it assumes a BFS has already been computed for the graph
func (g *Graph) Diameter() int {
	var max int
	for _, v := range g.V {
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
// Here the predecessor subgraph is:
// G{Predecessor} = (V, E{Predecessor})
func (g *Graph) DFS() int {

	// time is a global variable
	// Each vertex v has two timestamps, the first when it
	// is discovered (grayed) and the second when the search finishes
	// examining v's adjacency list (blackened). These timestamps
	// provide certain information about the structure of the graph
	// and are generally helpful in reasoning about the behaviour
	// of the depth-first search
	time := 0

	// could use a stack
	// for an iterative implementation
	// as opposed the current recursive implementation
	var DFSVisit func(*Vertex)

	DFSVisit = func(u *Vertex) {
		time++
		u.DStamp = time // discovered timestamp
		u.Color = Gray

		// visit every node reachable from each single
		// encountered node recursively
		node := u.adj.Head
		for node != nil {
			v, present := g.V[node.E]
			if present {
				if v.Color == White {
					v.Predecessor = u
					DFSVisit(v)
				}
			}
			node = node.Next
		}
		u.Color = Black
		time++
		u.FStamp = time // timestamp when dfs is done on this node

	}

	for _, u := range g.V {
		if u.Color == White {
			DFSVisit(u)
		}
	}
	return time
}

// TopSort performs a topological sort of the graph
// while by ordering all of its vertices linearly such that
// if graph G contains an edge (u, v), then u appears before v
// in the ordering - the graph must not contain cycles
// it receives as a parameter a reference to a linkedlist
// to be contain the list of ordered vertices
// the linkedList takes is of the string type - so the labels
// are the expected values to be added in the linear ordering
func (g *Graph) TopSort(ll *LinkedList) {
	time := 0
	var DFSVisit func(*Vertex)
	DFSVisit = func(u *Vertex) {
		time++
		u.DStamp = time
		u.Color = Gray
		node := u.adj.Head
		for node != nil {
			v, present := g.V[node.E]
			if present {
				if v.Color == White {
					v.Predecessor = u
					DFSVisit(v)
				}
			}
			node = node.Next
		}
		u.Color = Black
		time++
		u.FStamp = time
		ll.AddHead(u.Label)
	}
	for _, u := range g.V {
		if u.Color == White {
			DFSVisit(u)
		}
	}
}

// IsDAG performs a DFS on the graph and then returns true
// if the graph is a DAG ( Directed Acyclic Graph)
// a graph that has no back edge is a DAG.
// for edge (u, v) if u.d > v.d then there is a backedge
func (g *Graph) IsDAG() bool {
	isDag := true
	time := 0
	var DFSVisit func(*Vertex)
	DFSVisit = func(u *Vertex) {
		time++
		u.DStamp = time
		u.Color = Gray
		node := u.adj.Head
		for node != nil {
			v, present := g.V[node.E]
			if present {
				if v.Color == White {
					v.Predecessor = u
					DFSVisit(v)
				} else if v.DStamp > u.DStamp {
					isDag = false
					return
				}
			}
			node = node.Next
		}
		u.Color = Black
		time++
		u.FStamp = time
	}
	for _, u := range g.V {
		if u.Color == White {
			DFSVisit(u)
		}
	}
	return isDag
}

// DFSi is an iterative implementation of DFS
// using a stack returning the count for the last
// finishing time recorded for each visit order(θ(V+E)-time)
func (g *Graph) DFSi() int {
	time := 0
	var s *Vertex
	for _, k := range g.V {
		s = k
		break
	}
	s.Color = Gray
	time++
	s.DStamp = time

	stack := []*Vertex{}
	stack = append(stack, s)

	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		node := u.adj.Head
		for node != nil {
			v, present := g.V[node.E]
			if present {
				if v.Color == White {
					v.Predecessor = u
					v.Color = Gray
					time++
					v.DStamp = time
					stack = append(stack, v)
				}
			}
			node = node.Next
		}
		u.Color = Black
		time++
		u.FStamp = time
	}
	return time
}
