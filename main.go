package main

import (
	"fmt"

	"github.com/willpoint/algor/pkg/arithmetic"
)

func main() {
	i := arithmetic.RPMultiply(5, 5)
	j := arithmetic.AccMultiply(0, 5, 5)
	fmt.Println(i, j)
}
