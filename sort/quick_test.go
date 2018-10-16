package sort

import (
	"testing"
)

func TestQSort(t *testing.T) {

	ss := []byte("KJIHGFEDCBA")
	expected := "ABCDEFGHIJK"
	sorted := string(QSort(ss))
	if sorted != expected {
		t.Errorf("expected %s, got %s", expected, sorted)
	}

}
