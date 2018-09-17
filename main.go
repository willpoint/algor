package main

import "fmt"

func main() {

	a := []string{"a", "b"}
	fmt.Println(len(a))
	a = a[0 : len(a)-1]
	a = append(a, "m")
	fmt.Println(a)
}
