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

// BFS ...
func (g *Graph) BFS(s *Vertex) {
	Q := []*Vertex{}
	s.Color = Gray
	Q = append(Q, s)

	for len(Q) <= 0 {

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
