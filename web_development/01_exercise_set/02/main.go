// Initialize a MAP using a composite literal where the key is a string and the value is an int; print out the map; range over the map printing out just the key; range over the map printing out both the key and the value.
package main

import "fmt"

func main() {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for k := range m {
		fmt.Println(k)
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
