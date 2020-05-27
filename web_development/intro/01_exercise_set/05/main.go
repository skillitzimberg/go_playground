// Create a new type called “gator”. The underlying type of “gator” is an int. Using var, declare an identifier “g1” as type gator (var g1 gator). Assign a value to “g1”. Print out “g1”. Print the type of “g1” using fmt.Printf(“%T\n”, g1)

// Using var, declare an identifier “x” as type int (var x int). Print out “x”. Print the type of “x” using fmt.Printf(“%T\n”, x)

// Now add a method to type gator with this signature ...
// 	greeting()
// … and have it print “Hello, I am a gator”. Create a value of type gator. Call the greeting() method from that value.

// create another type called “flamingo”. Make the underlying type of “flamingo” bool. Give “flamingo” a method with this signature …
// greeting()
// … and have it print “Hello, I am pink and beautiful and wonderful.”

// create a new type “swampCreature”. The underlying type of “swapCreature” is interface. The behavior which the “swampCreature” interface defines is such that any type which has this method …
// 	greeting()
// … will implicitly implement the “swampCreature” interface. Create a func called “bayou” which takes a value of type “swampCreature” as an argument. Have this func print out the greeting for whatever “swampCreature” might be passed in.

package main

import "fmt"

type swampCreature interface {
	greeting()
}

func natureCalls(s swampCreature) {
	s.greeting()
}

type gator int

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

type flamingo bool

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful.")

}

func main() {
	var g1 gator
	g1 = 42
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x int
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// Q: Can you assign “g1” to “x”? Why or why not?
	// A: No. g1 is of type gator; x is of type int.
	// Note: It is possible to convert types that have the same underlying type.

	// Convert g1 (type gator) to type int
	x = int(g1)
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// Convert x (type int) to type gator
	x = 31
	g1 = gator(x)
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	g1.greeting()

	var f1 flamingo
	fmt.Println(f1)

	natureCalls(g1)
	natureCalls(f1)
}
