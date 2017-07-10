package main

import (
	"testing"
)

func BenchmarkFactorial(b *testing.B) {
	// run the factorial function b.N times
	for n := 0; n < b.N; n++ {
		main()
	}
}
