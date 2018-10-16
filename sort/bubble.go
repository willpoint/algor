package sort

// Bubble - θ(n2)-time) algorithm
func Bubble(A []byte) {
	for i := 0; i < len(A); i++ {
		for j := 1; j < len(A); j++ {
			if A[j-1] > A[j] {
				A[j-1], A[j] = A[j], A[j-1]
			}
		}
	}
}
