package main

import "fmt"

func main() {

	var m map[string]int

	_, ok := m["one"]
	if ok {
		fmt.Println("it exists")
	}
}
