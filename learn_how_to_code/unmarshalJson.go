package main

import (
	"encoding/json"
)

// How to make this generic, meaning how can I make this work with any targets of any type?
func unmarshalJSON(s string, target *[]interface{}) interface{} {
	err := json.Unmarshal([]byte(s), target)
	printError(err)
	return target
}
