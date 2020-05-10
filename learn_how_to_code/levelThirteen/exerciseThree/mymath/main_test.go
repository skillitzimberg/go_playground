package mymath

import "testing"

type test struct {
	xi []int
	f  float64
}

var tests = []test{
	{[]int{1, 4, 6, 8, 100}, 6},
	{[]int{0, 8, 10, 1000}, 9},
	{[]int{9000, 4, 10, 8, 6, 12}, 9},
	{[]int{123, 744, 140, 200}, 170},
}

func TestCenteredAvg(t *testing.T) {
	for _, test := range tests {
		expect := test.f
		got := CenteredAvg(test.xi)
		if expect != got {
			t.Errorf("Expected %v, got %v", expect, got)
		}
	}
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 4, 6, 8, 100})
	}
}
