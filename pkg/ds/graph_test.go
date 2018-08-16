package ds

import (
	"sync"
	"testing"
)

func TestGraph_BFS(t *testing.T) {
	type fields struct {
		Adj map[string]*Vertex
		mu  *sync.Mutex
	}
	type args struct {
		s *Vertex
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Adj: tt.fields.Adj,
				mu:  tt.fields.mu,
			}
			g.BFS(tt.args.s)
		})
	}
}
