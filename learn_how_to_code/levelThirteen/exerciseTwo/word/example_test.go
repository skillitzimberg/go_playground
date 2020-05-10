package word

import "fmt"

func ExampleCount() {
	fmt.Println(Count("Homer's journey ended with a surprise for Marge and Ned."))
	// Output:
	// 10
}

func ExampleUseCount() {
	fmt.Println(UseCount("She and I and her mom went to her mom and dad's."))
	// Output:
	// map[I:1 She:1 and:3 dad's.:1 her:2 mom:2 to:1 went:1]
}
