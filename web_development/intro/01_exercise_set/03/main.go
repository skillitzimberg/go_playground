// Create a new type called person which is STRUCT that stores fName and lName

// Use a composite literal to create a value of type person and assign it to a variable with the identifier “p1”; print out “p1”; print out just the field fName for “p1”.

// Take the STRUCT “person” in the previous exercise and add a field called “favFood” that stores a slice of string; for the variable “p1” use a composite literal to add values to the field favFood; print out the values in favFood; also print out the values for “favFood” by using a for range loop.

// Add a method to type “person” with the identifier “walk”. Have the func return this string: <person’s first name> is walking. Remember, you make a func a method by giving the func a receiver.
// 	func (r receiver) identifier(parameters) (returns) {
// 		<code>
// 	}
// To return a string, use fmt.Sprintln. Call the method assigning the returned string to a variable with the identifier “s”. Print out “s”.

package main

import "fmt"

type person struct {
	fName   string
	lName   string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fName, p.lName, "is walking")
}

func main() {

	p1 := person{
		"Guillermo",
		"Zimdars",
		[]string{"kibble", "chunks in gravy"},
	}

	walk := p1.walk()

	fmt.Println(p1)
	fmt.Println(p1.fName)
	fmt.Println(p1.favFood)
	for _, v := range p1.favFood {
		fmt.Println(v)
	}
	fmt.Println(walk)

}
