package sort

// QSort ...
func QSort(b []byte) []byte {
	qSort(b, 0, len(b)-1)
	return b
}

func qSort(b []byte, left, right int) {
	// ensure there are at least two items
	if left < right {
		pi := partition(b, left, right)
		qSort(b, left, pi-1)
		qSort(b, pi+1, right)
	}
}

func partition(b []byte, left, right int) int {
	// TODO: Use a different strategy to select a pivot
	// choosing left - is likely to result in the worst case
	// for this algorithm 0(n²) rather than the
	// average case of 0(n log n)
	pivot := left
	swap(b, pivot, right)
	store := left
	for i := left; i <= right-1; i++ {
		if b[i] <= b[right] {
			swap(b, i, store)
			store++
		}
	}
	swap(b, store, right)
	return store
}
