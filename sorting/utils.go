package sorting

func swap(b []byte, i, j int) {
	temp := b[i]
	b[i] = b[j]
	b[j] = temp
}
