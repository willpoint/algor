package graph

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

func TestBuildGraph(t *testing.T) {
	var param [][2]string
	f, err := os.Open("testdata/dgraph.json")
	if err != nil {
		t.Errorf("reading graph: %v", err)
	}
	dec := json.NewDecoder(f)
	err = dec.Decode(&param)
	if err != nil {
		t.Fatalf("decoding graph: %v", err)
	}
	G := BuildGraph(param)
	expected := uniqueLabels(param)
	if vnum := G.VNum; vnum != expected {
		t.Errorf("expected %d to equal %d, got %d", vnum, expected, vnum)
	}
	t.Log("dgraph:\n", G)
}

func TestBuildWeightedGraph(t *testing.T) {
	var param []struct {
		Pair   [2]string `json:"pair"`
		Weight int       `json:"weight"`
	}
	f, err := os.Open("testdata/wgraph.json")
	if err != nil {
		t.Errorf("reading graph: %v", err)
	}
	dec := json.NewDecoder(f)
	err = dec.Decode(&param)
	if err != nil {
		t.Fatalf("decoding graph: %v", err)
	}
	G := BuildWeightedGraph(param)
	var strings [][2]string
	for _, j := range param {
		strings = append(strings, j.Pair)
	}
	expected := uniqueLabels(strings)
	if vnum := G.VNum; vnum != expected {
		t.Errorf("expected %d to equal %d, got %d", vnum, expected, vnum)
	}
	t.Log("weighted graph:\n", G)
}

func TestTranspose(t *testing.T) {
	var param [][2]string
	f, err := os.Open("testdata/dgraph.json")
	if err != nil {
		t.Errorf("reading graph: %v", err)
	}
	dec := json.NewDecoder(f)
	err = dec.Decode(&param)
	if err != nil {
		t.Fatalf("decoding graph: %v", err)
	}
	G := BuildGraph(param)
	Gt := Transpose(G)
	if Gt.VNum != G.VNum {
		t.Error("G and Gt must have same vertex and edge lengths")
	}
	t.Log("transpose:\n", Gt)
}

// uniqueLabels in a graph = length of |V|
func uniqueLabels(ll [][2]string) int {
	d := map[string]int{}
	var sum int
	for i := 0; i < len(ll); i++ {
		for _, j := range ll[i] {
			if _, ok := d[j]; ok {
				continue
			}
			d[j]++
		}
	}
	for i := range d {
		sum += d[i]
	}
	return sum
}

func TestBFS(t *testing.T) {
	var param [][2]string
	f, err := os.Open("testdata/dgraph.json")
	if err != nil {
		t.Errorf("reading graph: %v", err)
	}
	dec := json.NewDecoder(f)
	err = dec.Decode(&param)
	if err != nil {
		t.Fatalf("decoding graph: %v", err)
	}
	G := BuildGraph(param)
	BFS(G, Label("A"))
	t.Log("bfs: \n", G)
}

func TestPrintPath(t *testing.T) {
	var param [][2]string
	f, err := os.Open("testdata/dgraph.json")
	if err != nil {
		t.Errorf("reading graph: %v", err)
	}
	dec := json.NewDecoder(f)
	err = dec.Decode(&param)
	if err != nil {
		t.Fatalf("decoding graph: %v", err)
	}
	G := BuildGraph(param)
	BFS(G, Label("A"))
	out := bytes.NewBuffer(make([]byte, 0, len(param)))
	PrintPath(out, G, Label("A"), Label("R"))
	t.Log("printpath:\n", out.String())
}

func TestDiameter(t *testing.T) {
	var param [][2]string
	f, err := os.Open("testdata/dgraph.json")
	if err != nil {
		t.Errorf("reading graph: %v", err)
	}
	dec := json.NewDecoder(f)
	err = dec.Decode(&param)
	if err != nil {
		t.Fatalf("decoding graph: %v", err)
	}
	G := BuildGraph(param)
	BFS(G, Label(param[0][0]))
	t.Log("diameter: \n", G.Diameter())
}

func TestDFS(t *testing.T) {
	var param [][2]string
	f, err := os.Open("testdata/dgraph.json")
	if err != nil {
		t.Errorf("reading graph: %v", err)
	}
	dec := json.NewDecoder(f)
	err = dec.Decode(&param)
	if err != nil {
		t.Fatalf("decoding graph: %v", err)
	}
	G := BuildGraph(param)
	DFS(G)
	t.Log("dfs recursive: \n", G)
}

func TestTopSort(t *testing.T) {
	var param [][2]string
	f, err := os.Open("testdata/complx.json")
	if err != nil {
		t.Errorf("reading graph: %v", err)
	}
	dec := json.NewDecoder(f)
	err = dec.Decode(&param)
	if err != nil {
		t.Fatalf("decoding graph: %v", err)
	}
	G := BuildGraph(param)
	l := TopoSort(G)
	t.Log("dress graph \n", G)
	t.Log("topological sort \n", l)
}

func TestDijsktra(t *testing.T) {
	var param []struct {
		Pair   [2]string `json:"pair"`
		Weight int       `json:"weight"`
	}
	f, err := os.Open("testdata/spgraph.json")
	if err != nil {
		t.Errorf("reading graph: %v", err)
	}
	dec := json.NewDecoder(f)
	err = dec.Decode(&param)
	if err != nil {
		t.Fatalf("decoding graph: %v", err)
	}
	G := BuildWeightedGraph(param)
	sp := Dijkstra(G, Label("s"))
	for _, v := range sp {
		t.Log(v)
	}
}
