package ds

import (
	"testing"
)

func TestGraph(t *testing.T) {

	a := []string{"abs", "jmh", "imo", "wil", "kha"}
	b := []string{"jmh", "imo", "abs"}
	c := []string{"imo", "abs", "kha", "wil"}
	d := []string{"dan", "wil", "imo"}
	e := []string{"wil", "jmh"}

	G, _ := NewGraph(a, b, c, d, e)

	tests := []struct {
		name string
		want int
	}{
		{"abs", 4},
		{"jmh", 2},
		{"imo", 3},
		{"dan", 2},
		{"wil", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			have := G.Adj[tt.name].VertexCount()
			if have != tt.want {
				t.Errorf("checking adj length - want %d, got %d", tt.want, have)
			}
		})
	}
}

func TestGraph_BFS(t *testing.T) {
	a := []string{"abs", "jmh", "imo", "wil", "kha"}
	b := []string{"jmh", "imo", "abs"}
	c := []string{"imo", "abs", "kha", "wil"}
	d := []string{"dan", "wil"}
	e := []string{"wil", "jmh"}
	f := []string{"kha", "min"}
	g := []string{"min"}

	G, _ := NewGraph(a, b, c, d, e, f, g)

	err := G.BFS("dan")
	if err != nil {
		t.Log("expected err to be non nil")
		t.Error(err.Error())
	}
	expect := 5
	if sp := G.Adj["min"].Distance; sp != expect {
		t.Errorf("expected %d, got, %d", expect, sp)
	}
}

func TestGraph_DFS(t *testing.T) {
	//
}
