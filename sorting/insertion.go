package sorting

// InsertionSort is achieved by spliting the
// given collection into two subcollections
// an already sorted collection and a yet to be
// sorted collection
// When the invariant is maintained that every
// new insertion into the sorted subcollection leaves
// it still sorted, the entire collection is then sorted after all
// elements from the unsorted subcollection are inserted
// into the sorted subcollection
// This algorithm works best for a small number of elements
// or when the elements it contain exist in a nearly sorted form
// an array already sorted in reverse order produces the worst
// performance for this algorithm.
// In time complexity - it's best case is 0(n)
// and it's average and worst case is 0(n²)
func InsertionSort(b []byte) []byte {
	for i := 1; i < len(b); i++ {
		insert(b, i, b[i])
	}
	return b
}

// insert inserts the element at the given position `pos`
// in the right position that keeps the sorted subcollection
// sorted, therefore maintaning the invariant
func insert(b []byte, pos int, val byte) {
	i := pos - 1
	for i >= 0 && b[i] > val {
		b[i+1] = b[i]
		i--
	}
	b[i+1] = val
}
