// Using the short declaration operator, create a variable with the identifier “s” and assign “i'm sorry dave i can't do that” to “s”.
//	 1. Print “s”.
//	 2. Print “s” converted to a slice of byte.
//	 3. Print “s” converted to a slice of byte and then converted back to a string.
//	 4. Using slicing, print just “i’m sorry dave”
//	 5. Using slicing, print just “dave i can’t”
//	 6. Using slicing, print just “can’t do that”
//	 7. print every letter of “s” with one rune (character) on each line
package main

import "fmt"

func main() {
	s := "i'm sorry dave i can't do that."
	fmt.Println(s)

	s1 := []byte(s)
	fmt.Println(s1)

	s2 := string(s1)
	fmt.Println(s2)

	s1a := string(s1[10:22])
	fmt.Println(s1a)

	s1b := string(s1[17:])
	fmt.Println(s1b)

	s1c := string(s1[:14])
	fmt.Println(s1c)

	for _, char := range s1 {
		fmt.Printf("%#U\n", char)
	}

}
