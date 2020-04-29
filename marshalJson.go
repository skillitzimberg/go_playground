package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func marshalJSON(input interface{}) {
	bs, err := json.Marshal(input)
	handleError(err)

	fmt.Println("Marshalled input written as a []byte:")
	fmt.Println(bs)
	fmt.Println()
	fmt.Println("Marshalled input written after conversion to a string:")
	fmt.Println(string(bs))
	fmt.Println()
	fmt.Println("Marshalled input written with os.Stdout.Write():")
	os.Stdout.Write(bs)
}
