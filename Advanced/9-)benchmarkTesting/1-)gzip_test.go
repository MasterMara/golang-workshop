package main

import "testing"

func BenchmarkGzipWithoutPoll(b *testing.B) {
	gzip := NewGzipWithoutPool()

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := gzip.Zip("Go Turkiye Egitim Kampi - 203 Test & Benchmarks")

		if err != nil {
			b.Fatal(err)
		}

	}
}

func BenchmarkGzipWithPoll(b *testing.B) {
	gzip := NewGzipWithPool()

	b.ResetTimer()   // Önceki Testten kalanları uçur
	b.ReportAllocs() // Memory Allocları ile ilgili bana bilgi ver
	for n := 0; n < b.N; n++ {
		err := gzip.Zip("Go Turkiye Egitim Kampi - 203 Test & Benchmarks")

		if err != nil {
			b.Fatal(err)
		}

	}
}

/* Notes By Mustafa Karacabey
1-) b.N --> Birim zamanda ben bunu bu kadar çalıştırabilirim. Miktar aslında
2-) go test -benchmem
3-) go test -bench . -cpu 2 (2 cpu ile ilgili benchmarkı yapar.)
4-) go test -bench . -benchtime 10x (ilgili n değeri 10 kez çalışacaktır.)
5-)  go test -bench . -benchtime 1s (1 saniye boyunca çalışacaktır.)
6-) go test -bench . -benchtime 3x -count 5 (5 kere aynı şey çalışacaktır.)
7-) go test -bench . gcflags='-N -l' (Compiler optimizasyonları kapatarak çalıştıracaktır.)
8-) bu tart pproof gibi memory ve cpu usageler ayrı şekilde öğrenilebilmektedir.
*/
