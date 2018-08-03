package sorting

// InsertionSort ...
func InsertionSort(b []byte) []byte {
	for i := 1; i < len(b); i++ {
		insert(b, i, b[i])
	}
	return b
}

func insert(b []byte, pos int, val byte) {
	i := pos - 1
	for i >= 0 && b[i] > val {
		b[i+1] = b[i]
		i--
	}
	b[i+1] = val
}
