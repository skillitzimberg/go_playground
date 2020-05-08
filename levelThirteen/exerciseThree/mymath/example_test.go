package mymath

import "fmt"

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 4, 6, 8, 100}))
	// Output:
	// 6
}
