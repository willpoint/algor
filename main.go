package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/willpoint/algor/graph"
)

func main() {
	var param [][2]string
	f, err := os.Open("sgraph.json")
	if err != nil {
		fmt.Printf("reading graph: %v", err)
	}
	dec := json.NewDecoder(f)
	err = dec.Decode(&param)
	G := graph.BuildGraph(param)
	fr := graph.DFS(G)
	fmt.Println("forests: ", fr)
	fmt.Println(G)

}
