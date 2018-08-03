package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/willpoint/algor/pkg/sorting"
)

func main() {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b1, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	start := time.Now()
	// _ = sorting.QSort(b1)
	_ = sorting.HeapSort(b1)
	done := time.Since(start)
	fmt.Println(done)

}
