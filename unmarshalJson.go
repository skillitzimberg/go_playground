package main

import (
	"encoding/json"
	"fmt"
)

// How to make this generic, meaning how can I make this work with any targets of any type?
func unmarshalJSONThingy(s string, target *[]thingy) {
	err := json.Unmarshal([]byte(s), target)
	printError(err)
	fmt.Println(target)
}
