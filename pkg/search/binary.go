package search

// Binary performs significantly better
// than a sequential search becuase it divides
// the collection|array in half until the target
// is found or not found.
// It's time complexity is 0(1) in the best case
// and 0(log n) in both average and worst cases
// This algorithm depends on having an already sorted
// data
func Binary(b []byte, t byte) bool {
	low := 0
	high := len(b) - 1

	for low <= high {
		i := (low + high) / 2
		if t == b[i] {
			return true
		}
		if t < b[i] {
			high = i - 1
		} else {
			low = i + 1
		}
	}
	return false
}
