package ds

import (
	"errors"
	"sync"
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
	Adj map[string]*Vertex
	mu  *sync.Mutex
}

// NewGraph ...
func NewGraph() *Graph {
	return &Graph{
		Adj: make(map[string]*Vertex),
	}
}

// Vertex ...
type Vertex struct {
	Label       string
	Color       Color
	Distance    int
	Predecessor *Vertex
	adj         map[string]*Vertex
	mu          *sync.Mutex
}

// Adjs ...
func (v *Vertex) Adjs() map[string]*Vertex {
	return v.adj
}

// AddEdge ...
func (v *Vertex) AddEdge(w *Vertex) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.adj[w.Label] = w
	w.adj[v.Label] = v
}

// NewVertex creates a new Vertex initialized with
// reasonable defaults
func NewVertex(l string) *Vertex {
	return &Vertex{
		Label: l,
		adj:   make(map[string]*Vertex),
		Color: White,
	}
}

// AddVertex ...
func (g *Graph) AddVertex(v *Vertex) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.Adj[v.Label]; ok {
		return ErrVertexExists
	}
	g.Adj[v.Label] = v
	return nil
}

// BFS, when given a source vertex s, systematically explores the
// edges of G to discover every vertex that is reachable from s
// It computes the smallest number of edges from s to each
// reachable vertex producing a `breadth-first-tree` with root s
// that contains all reachable vertices.
// To begin:
// It discovers all vertices at distance k from s before discovering
// vertices at distance k + 1
// It uses colors to keep track of the progress ensuring that
// each edge is visited only once.
// It uses a queue Q in this implementation to keep track of the
// traversal
func (g *Graph) BFS(s *Vertex) {
	s.Color = Gray

	Q := []*Vertex{}
	Q = append(Q, s)

	for len(Q) > 0 {

		u := Q[0]
		Q = Q[1:]

		for _, j := range u.Adjs() {
			if j.Color == White {
				j.Color = Gray
				j.Distance = u.Distance + 1
				j.Predecessor = u
				Q = append(Q, j)
			}
		}
		u.Color = Black
	}

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
func (g *Graph) DFS(s *Vertex) {
	//
}
