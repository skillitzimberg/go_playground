package dog

import (
	"testing"
)

type test struct {
	human int
	dog   int
}

var tests = []test{
	{10, 70},
	{2, 14},
	{30, 210},
	{49, 343},
	{-12, -84},
}

func TestYears(t *testing.T) {
	for _, v := range tests {
		hy := Years(v.human)
		if hy != v.dog {
			t.Errorf("Expected %d, but got %d", v.dog, hy)
		}

	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(20)
	}
}
