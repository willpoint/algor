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

	G := NewGraph(a, b, c, d, e)

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
			have := len(G.Adj[tt.name].Adjs())
			if have != tt.want {
				t.Errorf("checking adj length - want %d, got %d", have, tt.want)
			}
		})
	}

}
