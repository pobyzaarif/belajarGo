package thread

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumAlgA(t *testing.T) {
	sumA := SumAlgA(100)
	sumB := SumAlgA(100)
	sumC := SumAlgC(100)
	assert.Equal(t, sumA, sumB)
	assert.Equal(t, sumA, sumC)
}

func BenchmarkSumAlgA_100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAlgA(100)
	}
}

func BenchmarkSumAlgA_1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAlgA(1000)
	}
}

func BenchmarkSumAlgA_100000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAlgA(100000000)
	}
}

func BenchmarkSumAlgB_100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAlgB(100)
	}
}

func BenchmarkSumAlgB_1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAlgB(1000)
	}
}

func BenchmarkSumAlgB_100000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAlgB(100000000)
	}
}

func BenchmarkSumAlgC_100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAlgC(100)
	}
}

func BenchmarkSumAlgC_1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAlgC(1000)
	}
}

func BenchmarkSumAlgC_100000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAlgC(100000000)
	}
}
