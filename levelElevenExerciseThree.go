package main

import "fmt"

type myError struct {
	errorMessage string
}

func (me myError) Error() string {
	return fmt.Sprint(me.errorMessage)
}

func errorStructExercise() {
	err := myError{
		errorMessage: "My bad. Sorry.",
	}

	foo(err)

}
func foo(e error) {
	fmt.Println(e)
}
