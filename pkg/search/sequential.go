package search

// Sequential search is linear in nature.
// It searches through collection and returns true
// if the target element is found or false otherwise
// In the best, average and worst case -
// it's performance is 0(n)
// It can however be optimized if the underlying collection
// is to be search over again by moving the target element if
// found to the front of the collection - A[i, i -1] to A[1, i]
// making the next search 0(1). This serves as the basis for Most-Recently-Used
// paging algorithms.
func Sequential(b []byte, t byte) bool {
	for _, j := range b {
		if j == t {
			return true
		}
	}
	return false
}
