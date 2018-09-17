package arithmetic

import "testing"

func BenchmarkRPAMultiply(b *testing.B) {

	for i := 0; i < b.N; i++ {
		RPMultiply(41, 59)
	}

}

func BenchmarkAccMultiply(b *testing.B) {

	for i := 0; i < b.N; i++ {
		AccMultiply(0, 41, 59)
	}

}
func BenchmarkStrictAccMultiply(b *testing.B) {

	for i := 0; i < b.N; i++ {
		StrictAccMultiply(0, 41, 59)
	}

}

func BenchmarkIterativeAccMultiply(b *testing.B) {

	for i := 0; i < b.N; i++ {
		IterativeAccMultiply(0, 41, 59)
	}

}
