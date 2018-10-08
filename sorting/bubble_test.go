package sorting

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {

	ss := []byte("KJIHGFEDCBA")
	expected := "ABCDEFGHIJK"
	Bubble(ss)
	sorted := string(ss)
	if sorted != expected {
		t.Errorf("expected %s, got %s", expected, sorted)
	}

}
