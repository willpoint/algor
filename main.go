package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/willpoint/algor/graph"
)

func main() {

	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	var param [][2]string
	f, err := os.Open("scc.json")
	if err != nil {
		log.Fatal("reading graph: ", err)
	}

	dec := json.NewDecoder(f)
	err = dec.Decode(&param)
	if err != nil {
		log.Fatal("decoding graph: ", err)
	}

	G := graph.BuildGraph(param)
	graph.BFS(G, graph.Label("A"))
	time.Sleep(10 * time.Second)
	fmt.Println(G)
}
