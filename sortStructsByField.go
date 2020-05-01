package main

import (
	"fmt"
	"sort"
)

// ByName represents sorting by name with the Interface interface
type ByName []person

func (byName ByName) Len() int { return len(byName) }
func (byName ByName) Swap(i, j int) {
	byName[i], byName[j] = byName[j], byName[i]
}
func (byName ByName) Less(i, j int) bool { return byName[i].First < byName[j].First }

// ByAge represents sorting by name with the Interface interface
type ByAge []person

func (byAge ByAge) Len() int { return len(byAge) }
func (byAge ByAge) Swap(i, j int) {
	byAge[i], byAge[j] = byAge[j], byAge[i]
}
func (byAge ByAge) Less(i, j int) bool { return byAge[i].Age < byAge[j].Age }

func sortStructsByField() {
	u1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := person{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	people := []person{u1, u2, u3}

	// your code goes here
	sort.Stable(ByAge(people))
	for _, p := range people {
		fmt.Println(p.First, p.Last, p.Age)
		for _, s := range p.Sayings {
			fmt.Println("\t", s)

		}
		fmt.Println()
	}

	sort.Stable(ByName(people))
	for _, p := range people {
		fmt.Println(p.First, p.Last, p.Age)
		for _, s := range p.Sayings {
			fmt.Println("\t", s)

		}
		fmt.Println()
	}

}
