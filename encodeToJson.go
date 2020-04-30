package main

import (
	"encoding/json"
	"os"
)

func encodeJSON() {

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	// Unmarshal people from a string
	var people []person
	err := json.Unmarshal([]byte(s), &people)
	fakeHandleError(err)

	// Write the JSON encoding of a []person to the stream.
	encoder := json.NewEncoder(os.Stdout)

	err = encoder.Encode(people)
	fakeHandleError(err)
}
