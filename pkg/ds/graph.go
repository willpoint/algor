package ds

import (
	"errors"
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

// Graph ...
type Graph struct {
	Adj  map[string]*Vertex
	VNum int // number of vertices in Graph
}

// NewGraph is a variadic function that initializes a
// Graph G = (V, E) with a slice of slice of strings
// representing the label for each vertex v in G.
// The adjacency list for each vertex u ∈ V is represented
// as elements following string at first index of the slice S[1, 1i)
func NewGraph(ll ...[]string) (*Graph, error) {

	G := &Graph{
		Adj: make(map[string]*Vertex),
	}

	for _, j := range ll {
		v := NewVertex(j[0])
		for _, k := range j[1:] {
			v.adj.AddHead(k)
		}
		// each adjacency list is composed of
		// j[0] the item itself and j[1:]..., the elements
		// making up the adjlist for j[0]
		G.Adj[j[0]] = v
		G.VNum++
	}

	return G, nil
}

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

// Adjs ...
func (v *Vertex) Adjs() *LinkedList {
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

	s, present := g.Adj[label]
	if !present {
		return errors.New("vertex does not exist in graph")
	}

	s.Color = Gray

	Q := []*Vertex{}
	Q = append(Q, s)

	for len(Q) > 0 {

		u := Q[0] // pop the first vertex out of queue Q
		Q = Q[1:] // update queue to account for the popped vertex

		// u.adj is a list of adjacent vertices of u
		node := u.adj.Head
		for node != nil {
			v, present := g.Adj[node.E]
			if present {
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
			v, present := g.Adj[node.E]
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

	for _, u := range g.Adj {
		if u.Color == White {
			DFSVisit(u)
		}
	}
	return time
}
