package main

import "fmt"

func describePointers() {
	a := 9
	fmt.Println("a := 9")
	fmt.Println("The value of variable a is", a)
	fmt.Printf("a is of type %T\n", a)
	fmt.Println(`&a is a pointer. It points to the address in memory where the value of a is stored.
	The value (the memory address) returned by &a is `, &a)
	fmt.Printf(`&a is of type "pointer to an int" which is written as %T.
	"pointer to an int" means that the address being pointed to stores an int.`, &a)
	fmt.Println()

	b := &a
	fmt.Println("\nb := &a")
	fmt.Println("The value of b is ", b)
	fmt.Printf("b is of type %T\n, a pointer to an int.", b)
	fmt.Println("Value of &b: ", &b)
	fmt.Printf("&b is of type %T\n", &b)
	fmt.Println("Value stored at &b:", *b)
	fmt.Printf("*b is of type %T\n", *b)
}
