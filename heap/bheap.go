/*
Package heap provides a binary heap - a data structure that can be
viewed as a nearly complete binary tree represented in a slice.
implementing the Heaper interface which embeds the standard library
heap.Interface and few more methods is all it takes to use the package.
BHeap supports the two types of binary heap - max-heap and min-heap
the user has the responsibility of calling the right methods
on the BHeap type to maintain either a max-heap or min-heap property
Practical applications of a min-heap is for priority queues
while max-heap can be used by a heap-sort algorithm
*/
package heap

import (
	"container/heap"
)

// Heaper interface embeds the heap interface
// Pop() interface{}
// Push(x interface{})
// Len() int
// Less(i, j int) bool
// Swap(i, j int)
type Heaper interface {
	heap.Interface
	Get(i int) interface{}
	Set(i int, x interface{})
	Smaller(i int, key interface{}) bool
}

// BHeap structure
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

// MinHeapify maintains a min-heap property starting from element
// at node i. In a min-heap elem[parent(i)] <= elem[i] and the
// smallest element is at the root of the elem
func (b *BHeap) MinHeapify(i int) {
	smallest := i
	l, r := b.left(i), b.right(i)
	if l < b.HeapSize && b.Keys.Less(l, i) {
		smallest = l
	}
	if r < b.HeapSize && b.Keys.Less(r, smallest) {
		smallest = r
	}
	if smallest != i {
		b.Keys.Swap(smallest, i)
		b.MinHeapify(smallest)
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

// BuildMinHeap produces a min-heap for an unordered slice
func (b *BHeap) BuildMinHeap() {
	b.HeapSize = b.Keys.Len()
	mid := b.Keys.Len()/2 - 1
	for i := mid; i >= 0; i-- {
		b.MinHeapify(i)
	}
}

// HeapMaximum returns the maximum key in the heap
func (b *BHeap) HeapMaximum() interface{} {
	return b.Keys.Get(0)
}

// HeapMinimum return the minimum key in the heap
func (b *BHeap) HeapMinimum() interface{} {
	return b.Keys.Get(0)
}

// ExtractMax returns the maximum key in the heap and
// removes the element with the key
func (b *BHeap) ExtractMax() (interface{}, bool) {
	if b.HeapSize < 1 {
		return nil, false
	}
	max := b.Keys.Pop()
	b.HeapSize--
	b.MaxHeapify(0)
	return max, true
}

// ExtractMin returns the minimum key in the heap
// and removes the element with the key
func (b *BHeap) ExtractMin() (interface{}, bool) {
	if b.HeapSize < 1 {
		return nil, false
	}
	min := b.Keys.Pop()
	b.HeapSize--
	b.MinHeapify(0)
	return min, true
}

// HeapIncreaseKey increases the key at index i provided
// the key is greater than the existing one
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

// HeapDecreasekey is an inverse of HeapIncreaseKey
func (b *BHeap) HeapDecreasekey(i int, key interface{}) bool {
	if !b.Keys.Smaller(i, key) {
		return false
	}
	b.Keys.Set(i, key)
	for i > 0 && b.Keys.Less(i, b.parent(i)) {
		b.Keys.Swap(i, b.parent(i))
		i = b.parent(i)
	}
	return true
}

// MaxHeapInsert inserts a new element into the heap with
// the given key
func (b *BHeap) MaxHeapInsert(key interface{}) {
	b.HeapSize++
	b.Keys.Push(key)
	b.HeapIncreaseKey(b.HeapSize-1, key)
}

// Sort sorts the heap in non decreasing order in place
// it uses a max-heap in its implementation
func (b *BHeap) Sort() {
	b.BuildMaxHeap()
	for i := b.Keys.Len() - 1; i > 0; i-- {
		b.MaxHeapify(0)
		b.Keys.Swap(0, i)
		b.HeapSize--
	}
}
