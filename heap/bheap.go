package heap

// BHeap is a binary heap - a data structure that can be
// viewed as a nearly complete binary tree with each
// node representing an element in the slice
// The root of the binary heap is at elem[0] and given
// a node i, we can compute the indices of its parent, left
// child and right child
// In the two kinds of binary heap which this type can be used
// for, the user has the responsibility of calling the right methods
// on the BHeap type to maintain either a max-heap or min-heap property
// for either a max-heap or a min-heap
// while the reverse is the case for a min-heap
// elem[parent(i)] <= elem[i], with smallest element at root
// Practical applications of a min-heap is for priority queues
// while max-heap can be used by a heap-sort algorithm
import (
	"container/heap"
)

// Heaper interface implements the heap interface
// Pop() interface{}
// Push(x interface{})
// Len() int
// Less(i, j int) bool
// Swap(i, j int)
// Get(i int) interface{}
// Set(i int, x interface{})
type Heaper interface {
	heap.Interface
	Get(i int) interface{}
	Set(i int, x interface{})
	Smaller(i int, key interface{}) bool
}

// BHeap ... ...
type BHeap struct {
	Length   int
	HeapSize int
	Keys     Heaper
}

// NewBinaryHeap initializes a binary heap
func NewBinaryHeap(h Heaper) *BHeap {
	return &BHeap{
		HeapSize: h.Len(),
		Keys:     h,
	}
}

// parent returns the index of an element with at index i
func (b *BHeap) parent(i int) int {
	return i >> 1 // produces the floor
}

// left return index of left child of element at index i
func (b *BHeap) left(i int) int {
	return i<<1 + 1
}

// right return index of right child of element at index i
func (b *BHeap) right(i int) int {
	return i<<1 + 2
}

// MaxHeapify maintains a max-heap property starting from the
// given index i. In a max-heap
// elem[parent(i)] >= elem[i] and the largest
// element is stored at the root, and the subtree rooted at node i
// contains no items larger than that contained at the node itself.
func (b *BHeap) MaxHeapify(i int) {
	largest := i
	l, r := b.left(i), b.right(i)
	if l < b.HeapSize && b.Keys.Less(i, l) {
		largest = l
	}
	if r < b.HeapSize && b.Keys.Less(largest, r) {
		largest = r
	}
	if largest != i {
		b.Keys.Swap(largest, i)
		b.MaxHeapify(largest)
	}
}

// BuildMaxHeap produces a max-heap for an unordered slice
func (b *BHeap) BuildMaxHeap() {
	b.HeapSize = b.Keys.Len()
	mid := b.Keys.Len()/2 - 1
	for i := mid; i >= 0; i-- {
		b.MaxHeapify(i)
	}
}

// HeapMaximum returns the maximum key in the heap
func (b *BHeap) HeapMaximum() interface{} {
	return b.Keys.Get(0)
}

// ExtractMax ...
func (b *BHeap) ExtractMax() (interface{}, bool) {
	if b.HeapSize < 1 {
		return nil, false
	}
	max := b.Keys.Pop()
	b.HeapSize--
	b.MaxHeapify(0)
	return max, true
}

// HeapIncreaseKey ...
func (b *BHeap) HeapIncreaseKey(i int, key interface{}) bool {
	if b.Keys.Smaller(i, key) {
		return false
	}
	b.Keys.Set(i, key)
	for i > 0 && b.Keys.Less(b.parent(i), i) {
		b.Keys.Swap(i, b.parent(i))
		i = b.parent(i)
	}
	return true
}

// MaxHeapInsert ...
func (b *BHeap) MaxHeapInsert(key interface{}) {
	b.HeapSize++
	b.Keys.Push(key)
	b.HeapIncreaseKey(b.HeapSize-1, key)
}
