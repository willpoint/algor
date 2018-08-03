package sorting

// HeapSort attempts to reduce the number of comparisons that need to be
// made when sorting an array by utilizing properties of a binary tree.
// Shape property - every node at depth K > 0 can exist only if 2^k-1 nodes at
// depth (k -1) exists and that nodes are filled from `left to right`
// Heap property - that every node in the tree contains a value greater than or
// equal to either of its two children, if it has any.
// This algorithm first structures the given array into that of a heap
// For a given array - the root is labeled 0. For a node with label i, its
// left child is labeled 2*i + 1 & its right child is labeled 2*i + 2.
// For a non-root node labeled i, its parent node is (floor)(i - 1)/2
func HeapSort(b []byte) []byte {
	// turn b into a heap structured array
	buildHeap(b)
	for i := len(b) - 1; i > 0; i-- {
		swap(b, 0, i)
		heapify(b, 0, i)
	}
	return b
}

func buildHeap(b []byte) {
	mid := int(len(b)/2) - 1
	for i := mid; i >= 0; i-- {
		heapify(b, i, len(b))
	}
}

func heapify(b []byte, idx, max int) {
	left := 2*idx + 1
	right := 2*idx + 2
	var largest int
	if left < max && b[left] > b[idx] {
		largest = left
	} else {
		largest = right
	}
	if right < max && b[right] > b[largest] {
		largest = right
	}
	if largest != idx {
		swap(b, idx, largest)
	}
	heapify(b, largest, max)
}
