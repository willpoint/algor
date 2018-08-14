package ds

import (
	"errors"
)

// The choice to represent a graph in this package
// is the adjacency list |A|(adj) as it provides a compact
// way to represent sparse graphs - for which |E| edge is much
// less than |v2|
// The adjacency-list representation of a graph G = (V, E)
// consists of a slice adj of |V| lists, one for each
// vertex in V. For each u ∈ V, the adjacency list Adj[u]
// contains all the vertices v such that there is an edge
// (u, v) ∈ E. That is, Adj[u] consists of all the vertices
// adjacent to u in G or pointers to these vertices.
// Since the adjacency list represents the edges of a graph,
// we treat the slice Adj as an attribute of the graph
// just as we treat the edge set E. Therefore G.Adj[u]

var (
	ErrNotIncidentVertex = errors.New("vertex not incident on edge")
)

// Graph ...
type Graph struct {
	Vertices []*Vertex
	Edges    []*Edge
}

// InsertVertex ...
func (g *Graph) InsertVertex(l string) *Vertex {
	return nil
}

// InsertEdges ...
func (g *Graph) InsertEdges(v, w *Vertex, l string) *Edge {
	e := &Edge{
		Label:       l,
		EndVertices: [2]*Vertex{v, w},
	}
	g.Edges = append(g.Edges, e)
	return e
}

// EraseVertex ...
func (g *Graph) EraseVertex(v *Vertex) {
	// remove vertex v and all its incident edges
	index := 0
	for _, j := range g.Vertices {
		index++
		if j.Label == v.Label {
			break
		}
	}
	pV := g.Vertices[:index-1]
	nV := g.Vertices[index:]
	g.Vertices = append(pV, nV...)

}

// Vertex ...
type Vertex struct {
	Label         string
	IncidentEdges []*Edge
}

// IsAdjacentTo ...
func (v *Vertex) IsAdjacentTo(u *Vertex) bool {
	for _, j := range v.IncidentEdges {
		if j.EndVertices[0].Label == u.Label ||
			j.EndVertices[1].Label == u.Label {
			return true
		}
	}
	return false
}

// Edge ...
type Edge struct {
	Label       string
	EndVertices [2]*Vertex
}

// Opposite returns the end vertex of edge e distinct from
// vertex v, and an error if e is not incident on v
func (e *Edge) Opposite(u *Vertex) (*Vertex, error) {
	switch {
	case e.EndVertices[0].Label == u.Label:
		{
			return e.EndVertices[1], nil
		}
	case e.EndVertices[1].Label == u.Label:
		{
			return e.EndVertices[0], nil
		}
	default:
		return nil, ErrNotIncidentVertex
	}
}

// IsAdjacentTo tests if edges e and f are adjacent
func (e *Edge) IsAdjacentTo(f *Edge) bool {
	for _, j := range e.EndVertices {
		if j.Label == f.EndVertices[0].Label ||
			j.Label == f.EndVertices[1].Label {
			return true
		}
	}
	return false
}
