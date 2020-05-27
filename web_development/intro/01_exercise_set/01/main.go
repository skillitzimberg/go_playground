// Initialize a SLICE of int using a composite literal; print out the slice; range over the slice printing out just the index; range over the slice printing out both the index and the value
package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range xi {
		fmt.Println(i)
	}
	for i, v := range xi {
		fmt.Println(i, v)
	}
}
