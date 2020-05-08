package main

import (
	"fmt"

	"github.com/skillitzimberg/go_playground/levelThirteen/exerciseTwo/word"
)

func main() {
	s := "hasdfi iaisdfn ijknidf iinji ij"
	fmt.Println(s)
	c := word.Count(s)
	fmt.Println(c)
}
