package sort

import (
	"testing"
)

func TestHeapSort(t *testing.T) {

	ss := []byte("KJIHGFEDCBA")
	expected := "ABCDEFGHIJK"
	sorted := string(HeapSort(ss))
	if sorted != expected {
		t.Errorf("expected %s, got %s", expected, sorted)
	}

}
