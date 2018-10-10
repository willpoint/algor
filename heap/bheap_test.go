package heap

import "testing"

type elem []int

func (e elem) Less(i, j int) bool {
	return e[i] < e[j]
}

func (e *elem) Swap(i, j int) {
	ee := *e
	ee[i], ee[j] = ee[j], ee[i]
}

func (e elem) Len() int {
	return len(e)
}

func (e *elem) Pop() interface{} {
	k := *e
	o, j := k[0], k[1:]
	*e = j
	return o
}

func (e *elem) Push(x interface{}) {
	k := *e
	k = append(k, x.(int))
	*e = k
}

func (e elem) Smaller(i int, x interface{}) bool {
	return x.(int) < e.Get(i).(int)
}

func (e elem) Get(i int) interface{} {
	return e[i]
}

func (e *elem) Set(i int, x interface{}) {
	k := *e
	k[i] = x.(int)
}

func TestMinHeap(t *testing.T) {
	var e elem = []int{23, 77, 5, 7, 8, 107, 3, 11, 13}
	h := NewBinaryHeap(&e)
	h.BuildMinHeap()
	prevLen := len(e)
	smallest := 3
	if min, _ := h.ExtractMin(); min != smallest {
		t.Errorf("expected %d to be %d, got %d", min, smallest, min)
	}
	if prevLen != h.Keys.Len()+1 {
		t.Errorf("expected length of elem to decrease by one")
	}

}

func TestHeapSort(t *testing.T) {
	var e elem = []int{23, 77, 5, 7, 8, 107, 3, 11, 13}
	h := NewBinaryHeap(&e)
	h.Sort()
	largest := 107
	result := e[len(e)-1]
	t.Log(h.Keys)
	if result != largest {
		t.Errorf("expected %d to be %d, got %d", result, largest, result)
	}
}
