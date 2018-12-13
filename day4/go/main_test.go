package main

import "testing"

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strategy1()
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strategy2()
	}
}
