package main

import "fmt"

type triple struct {
	a, b, c int
}

// mod is equivalent to a mod b - but compensates for
// when a is negative
func mod(a, b int) int {
	i := a % b
	// if a is negative
	if a/-1 >= 0 {
		return b + i
	}
	return i
}

// find the number that is relatively prime to another
// euclid is a step there
// I need to find a small odd integer e that is relatively
// prime to r. Two numbers are relatively prime if their gcd
// greatest common divisor is 1.
// Using the theorem in number theory that if a and b
// are integers, not both zero, then their  gcd
// g equals ai + bj for some integer i and j
func euclid(a, b int) triple {
	t := triple{}
	if b == 0 {
		t.a = a
		t.b = 1
		t.c = 0
		return t
	}

	// if euclid(r, e) is called
	// the number of recursive calls is 0(log e)
	i := euclid(b, mod(a, b))
	t.a = i.a
	t.b = i.c
	t.c = i.b - (a/b)*i.c
	return t
}

func main() {
	fmt.Println(euclid(18, 30))
}
