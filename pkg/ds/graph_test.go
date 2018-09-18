package ds

import (
	"testing"
)

func TestGraph(t *testing.T) {

	a := []string{"abs", "jmh", "imo", "wil"}
	b := []string{"jmh", "imo", "abs"}
	c := []string{"imo", "abs", "kha", "wil"}
	d := []string{"dan", "wil", "imo"}
	e := []string{"wil", "jmh"}

	G, _ := BuildGraph(a, b, c, d, e)

	tests := []struct {
		name string
		want int
	}{
		{"abs", 3},
		{"jmh", 2},
		{"imo", 3},
		{"dan", 2},
		{"wil", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			have := G.V[tt.name].VertexCount()
			if have != tt.want {
				t.Errorf("checking adj length - want %d, got %d", tt.want, have)
			}
		})
	}

	// test correctness of edge creation
	edgeNum := 11 // by manual count (test graph is undirected)
	if got := G.ENum; got != edgeNum {
		t.Errorf("checking edge count - want %d, got %d", edgeNum, got)
	}
}

func TestGraph_BFS(t *testing.T) {
	a := []string{"abs", "jmh", "imo", "wil", "kha"}
	b := []string{"jmh", "imo", "abs"}
	c := []string{"imo", "abs", "kha", "wil"}
	d := []string{"dan", "wil"}
	e := []string{"wil", "jmh"}
	f := []string{"kha", "min"}
	g := []string{"min", "kha"}

	G, _ := BuildGraph(a, b, c, d, e, f, g)

	err := G.BFS("dan")
	if err != nil {
		t.Log("expected err to be non nil")
		t.Error(err.Error())
	}
	expect := 5
	if sp := G.V["min"].Distance; sp != expect {
		t.Errorf("expected %d, got, %d", expect, sp)
	}
}

func TestGraph_DFS(t *testing.T) {
	a := []string{"abs", "jmh", "imo", "wil", "kha"}
	b := []string{"jmh", "imo", "abs"}
	c := []string{"imo", "abs", "kha", "wil"}
	d := []string{"dan", "wil"}
	e := []string{"wil", "jmh"}
	f := []string{"kha", "dan"}

	G, _ := BuildGraph(a, b, c, d, e, f)

	// There are 6 vertices in the generated graph
	// given that time t is incremented twice for each vertex
	// once for when it is discovered and once for when
	// a dfs search has been completed for that vertex
	// 6 * 2 is the final value of the time
	expected := G.VNum * 2
	if got := G.DFS(); got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestGraph_DFSi(t *testing.T) {
	a := []string{"abs", "jmh", "imo", "wil", "kha"}
	b := []string{"jmh", "imo", "abs"}
	c := []string{"imo", "abs", "kha", "wil"}
	d := []string{"dan", "wil"}
	e := []string{"wil", "jmh"}
	f := []string{"kha", "dan", "wil"}

	G, _ := BuildGraph(a, b, c, d, e, f)

	// There are 6 vertices in the generated graph
	// given that time t is incremented twice for each vertex
	// once for when it is discovered and once for when
	// a dfs search has been completed for that vertex
	// 6 * 2 is the final value of the time
	expected := G.VNum * 2
	if got := G.DFSi(); got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

// TestGraph_GraphTranspose tests that each outgoing vertex u -> v
// is reversed to result in v -> u, the test therefore checks
// that G and Gt maintain the same properties but for the new directions
func TestGraph_GraphTranspose(t *testing.T) {
	a := []string{"abs", "jmh", "imo", "wil", "kha"}
	b := []string{"jmh", "imo", "abs"}
	c := []string{"imo", "abs", "kha", "wil"}
	d := []string{"dan", "wil"}
	e := []string{"wil", "jmh"}
	f := []string{"kha", "dan", "wil"}

	G, _ := BuildGraph(a, b, c, d, e, f)
	Gt := GraphTranspose(G)
	if av, bv := G.VNum, Gt.VNum; av != bv {
		t.Errorf("expected %d to be equal to %d", av, bv)
	}
	if ae, be := G.ENum, Gt.ENum; ae != be {
		t.Errorf("expected %d to be equal to %d", ae, be)
	}
}

func TestGraph_TopSort(t *testing.T) {

	a := []string{"undershorts", "shoes", "pants"}
	b := []string{"pants", "belt", "shoes"}
	c := []string{"belt", "jacket"}
	d := []string{"jacket"}
	e := []string{"tie", "jacket"}
	f := []string{"shirt", "belt", "tie"}
	g := []string{"socks", "shoes"}
	h := []string{"watch"}
	i := []string{"shoes"}
	G, err := BuildGraph(a, b, c, d, e, f, g, h, i)
	if err != nil {
		t.Error("expected nil error", err)
	}
	ll := NewLinkedList()
	// Do TopSort
	G.TopSort(ll)
	ss := []string{}
	node := ll.Head
	for node != nil {
		ss = append(ss, node.E)
		node = node.Next
	}
	pre, post := "undershorts", "pants"
	if prec := preceeds(pre, post, ss); prec != true {
		t.Errorf("Expected %s to preceed %s", pre, post)
	}
}

// preceeds tells if string a is seen before string b
// in a given slice
func preceeds(a, b string, ss []string) bool {
	var aindex int
	for i, s := range ss {
		if a == s {
			aindex = i
			break
		}
	}
	for _, s := range ss[aindex:] {
		if s == b {
			return true
		}
	}
	return false
}
