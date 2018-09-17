package arithmetic

// odd returns true if n is odd
// by comparing the value resulting
// from and(ing) to least significant
// bit of n with 0x1 to 1
func odd(n int) bool {
	return n&0x1 == 1
}

// half for this algorithm implies that:
// if a number is odd, its halving the number
// results in half(n) == half(n-1)
// if 3 is the number then half(3) == half(3-1)
func half(n int) int {
	return n >> 1
}

// LinearMultiply provides an 0(n) time
// complexity for values of n
func LinearMultiply(n, a int) int {
	if n == 1 {
		return a
	}
	return LinearMultiply(n-1, a) + a
}

// RPMultiply uses the Russian Peasant Algorithm
// described in Alexandar Stepanov's `From Mathematics to Generic Programming`
// providing an order 0(log n) time complexity for all values of n
func RPMultiply(n, a int) int {
	if n == 1 {
		return a
	}
	r := RPMultiply(half(n), a+a)
	if odd(n) {
		r = r + a
	}
	return r
}

// AccMultiply uses r as a running result that
// accumulates the partial products n*a
// making use of a tail recursion
func AccMultiply(r, n, a int) int {
	if odd(n) {
		r = r + a
		if n == 1 {
			return r
		}
	}
	return AccMultiply(r, half(n), a+a)
}

// StrictAccMultiply uses r as a running result that
// accumulates the partial products n*a
// making use of a strict tail recursion
// defined as a ruecursion where all the tail-recursive
// calls are done with the formal parameters of the procedure being
// the corresponding arguments
func StrictAccMultiply(r, n, a int) int {
	if odd(n) {
		r = r + a
		if n == 1 {
			return r
		}
	}
	n = half(n)
	a = a + a
	return StrictAccMultiply(r, n, a)
}

// IterativeAccMultiply uses r as a running result that
// accumulates the partial products n*a iteratively
func IterativeAccMultiply(r, n, a int) int {
	for {
		if odd(n) {
			r = r + a
			if n == 1 {
				return r
			}
		}
		n = half(n)
		a = a + a
	}
}
