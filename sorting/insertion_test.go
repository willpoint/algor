package sorting

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {

	ss := []byte("KJIHGFEDCBA")
	expected := "ABCDEFGHIJK"
	sorted := string(InsertionSort(ss))
	if sorted != expected {
		t.Errorf("expected %s, got %s", expected, sorted)
	}

}
