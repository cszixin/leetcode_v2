package pkg

import "testing"

func BenchmarkFibv1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fibv1(40)

	}

}

func BenchmarkFibv2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fibv2(40)

	}

}

func BenchmarkFibv3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fibv3(40)

	}

}

func BenchmarkFibv4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fibv4(40)

	}

}
