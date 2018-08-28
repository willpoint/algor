package main

import (
	"fmt"

	"github.com/willpoint/algor/pkg/ds"
)

func main() {

	// m := ds.NewMatrix(2, 3)
	// n := ds.NewMatrix(3, 2)

	// _ = m.SetRowSlice(0, []float64{-1, 2, 0})
	// _ = m.SetRowSlice(1, []float64{4, 1, -2})

	// _ = n.SetRowSlice(0, []float64{1, 3})
	// _ = n.SetRowSlice(1, []float64{0, -1})
	// _ = n.SetRowSlice(2, []float64{7, 5})

	// o, _ := m.Multiply(n)
	// j := o.Transpose()
	// fmt.Printf("%#v\n", j.Array())

	// t := ds.NewMatrix(2, 3)
	// _ = t.SetRowSlice(0, []float64{2, 4, -1})
	// _ = t.SetRowSlice(1, []float64{0, 3, 5})
	// t1 := t.Transpose()
	// fmt.Printf("%#v\n", t1.Array())

	// z := ds.Identity(3)
	// fmt.Printf("%#v\n", z.Array())

	q := ds.NewMatrix(3, 3)
	_ = q.SetRowSlice(0, []float64{-1, 2, 0})
	_ = q.SetRowSlice(1, []float64{4, 1, -2})
	_ = q.SetRowSlice(2, []float64{8, 6, -1})

	r, _ := q.ShedRow(0)
	fmt.Printf("%#v\n", r.Array())

}
