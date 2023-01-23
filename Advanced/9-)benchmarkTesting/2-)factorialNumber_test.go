package main

import "testing"

func FindFactorial(number int) int {
	if number == 0 {
		return 1
	}
	return FindFactorial(number-1) * number
}

func BenchmarkFindFactorial(b *testing.B) {
	//b.N --> bu değer Benchmark Çalıştırıcısı tarafından belirlenecektir.

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		FindFactorial(15)
	}

}
