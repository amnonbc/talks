package llist

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var savedSum int


func TestAdd(t *testing.T) {
	ll := List{}
	ll.Insert(3)
	ll.Insert(1)
	ll.Insert(2)
	assert.Equal(t, "1 2 3 ", ll.String())
}

func TestSum(t *testing.T) {
	ll := List{}
	ll.Insert(3)
	ll.Insert(1)
	ll.Insert(2)
	assert.Equal(t, 6, ll.Sum())
}


func benchmarkSum(num int, b *testing.B) {
	ll := List{}
	ll.Fill(num)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		savedSum = ll.Sum()
	}
}

func benchmarkSumV(num int, b *testing.B) {
	v := make([]int, num, num)
	fillV(v)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		savedSum = sumV(v)
	}
}

// STARTBENCH OMIT
func BenchmarkSumV10(b *testing.B)  { benchmarkSumV(10, b) }
func BenchmarkSumV100(b *testing.B)  { benchmarkSumV(100, b) }
func BenchmarkSumV1000(b *testing.B)  { benchmarkSumV(1000, b) }
func BenchmarkSumV10000(b *testing.B)  { benchmarkSumV(10000, b) }
func BenchmarkSumV100000(b *testing.B)  { benchmarkSumV(100000, b) }
// ENDBENCH OMIT

func BenchmarkSum10(b *testing.B)  { benchmarkSum(10, b) }
func BenchmarkSum100(b *testing.B)  { benchmarkSum(100, b) }
func BenchmarkSum1000(b *testing.B)  { benchmarkSum(1000, b) }
func BenchmarkSum10000(b *testing.B)  { benchmarkSum(10000, b) }
func BenchmarkSum100000(b *testing.B)  { benchmarkSum(100000, b) }
