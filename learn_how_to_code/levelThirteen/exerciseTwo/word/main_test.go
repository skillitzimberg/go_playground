package word

import (
	"testing"
)

type countTest struct {
	s string
	c int
}

var countTests = []countTest{
	{"frank", 1},
	{"frank done finished", 3},
	{"Thorough      ", 1},
	{"      poor andy", 2},
}

type useCountTest struct {
	s string
	m map[string]int
}

var useCountTests = []useCountTest{
	{"frank", map[string]int{"frank": 1}},
	{"frank done finished done frank", map[string]int{"frank": 2, "done": 2, "finished": 1}},
	{"Thorough      ", map[string]int{"Thorough": 1}},
	{"      poor andy", map[string]int{"poor": 1, "andy": 1}},
}

func TestCount(t *testing.T) {
	for _, v := range countTests {
		expect := v.c
		got := Count(v.s)

		if expect != got {
			t.Errorf("Expected %v, got %v", expect, got)
		}
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("Don't want to be your beast of burden")
	}
}

func TestUseCount(t *testing.T) {
	for _, v := range useCountTests {
		for str := range v.m {
			expect := v.m[str]
			got := UseCount(v.s)[str]
			if expect != got {
				t.Errorf("Expected %v, got %v", expect, got)
			}
		}
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("Don't want to be your beast of burden")
	}
}
