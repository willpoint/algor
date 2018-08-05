package main

import (
	"fmt"

	"github.com/willpoint/algor/pkg/search"
	_ "github.com/willpoint/algor/pkg/sorting"
)

func main() {
	// file, err := os.Open("words.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer file.Close()

	// b1, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// start := time.Now()
	// // st := sorting.QSort(b1)
	// st := sorting.HeapSort(b1)
	// // st := sorting.InsertionSort(b1)
	// done := time.Since(start)
	// fmt.Println(string(st))
	// fmt.Println(done)
	l := search.Hash([]string{"hypoplankton"}, "man")
	fmt.Println(l)

}
