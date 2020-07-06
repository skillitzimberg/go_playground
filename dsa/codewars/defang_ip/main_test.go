package main

import "testing"

type test struct {
	ip      string
	isValid bool
}

var tests = []test{
	{"1.2.3.4", true},
	{"123.45.67.89", true},
	{"192.168.1.300", true},
	{"0.34.82.53", true},
	{"0.0.0.0", true},
	{"127.1.1.0", true},
	{"1.2.3", false},
	{"1.2.3.4.5", false},
	{"123.456.78.90", false},
	{"123.045.067.089", false},
}

func TestIsValidIP(t *testing.T) {
	for _, v := range tests {
		expect := v.isValid
		got := isValidIP(v.ip)
		if got != expect {
			t.Errorf("Expected %v, but got %v", expect, got)
		}
	}
}

func BenchmarkIsValidIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValidIP("192.168.1.300")
	}
}
