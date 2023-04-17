package popcount

import "testing"

const population uint64 = 0x1234567890ABCDEF

func BenchmarkPopCountTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTable(population)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(population)
	}
}

func BenchmarkPopBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopBit(population)
	}
}

func BenchmarkPopClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopClear(population)
	}
}
