package thread

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumAlgA(t *testing.T) {
	sumA := SumAlgA(1000)
	sumB := SumAlgA(1000)
	sumC := SumAlgC(1000)
	assert.Equal(t, sumA, sumB)
	assert.Equal(t, sumA, sumC)
}

func BenchmarkSumAlgA_10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAlgA(10)
	}
}

func BenchmarkSumAlgA_100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAlgA(100)
	}
}

func BenchmarkSumAlgA_1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAlgA(10000000)
	}
}

func BenchmarkSumAlgB_10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAlgB(10)
	}
}

func BenchmarkSumAlgB_100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAlgB(100)
	}
}

func BenchmarkSumAlgB_1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAlgB(10000000)
	}
}

func BenchmarkSumAlgC_10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAlgC(10)
	}
}

func BenchmarkSumAlgC_100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAlgC(100)
	}
}

func BenchmarkSumAlgC_1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAlgC(10000000)
	}
}
