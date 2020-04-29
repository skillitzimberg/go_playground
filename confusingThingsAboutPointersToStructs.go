package main

import "fmt"

type thing struct {
	name string
}

func confusingThingsAboutPointersToStructs() {

	a := 43

	c := thing{
		name: "Booger",
	}

	fmt.Printf("Value of a: %v; \nType of a: %T; \nValue of pointer to a (the address of a): %p; \nType of pointer to a: %T\n", a, a, &a, &a)
	fmt.Println()

	fmt.Printf("Value of c: %v; \nType of c: %T; \nValue of pointer to c (the address of c): %p; \nType of pointer to c: %T\n", c, c, &c, &c)
	fmt.Println()

	var d = &c
	fmt.Printf("Value of d: %v; \nType of d: %T; \nValue of pointer to d (the address of d): %p; \nType of pointer to d: %T\n", d, d, &d, &d)
	fmt.Println()

	var e = c
	fmt.Printf("Value of e: %v; \nType of e: %T; \nValue of pointer to e (the address of e): %p; \nType of pointer to e: %T\n", e, e, &*&e, &e)
	fmt.Println()

	d.name = "Sugar"
	fmt.Printf("Value of d: %v; \nType of d: %T; \nValue of pointer to d (the address of d): %p; \nType of pointer to d: %T\n", d, d, &d, &d)
	fmt.Println()

	fmt.Printf("Value of c: %v; \nType of c: %T; \nValue of pointer to c (the address of c): %p; \nType of pointer to c: %T\n", c, c, &c, &c)
	fmt.Println()

}
