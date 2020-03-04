package main

import "fmt"

type person struct {
	first            string
	last             string
	favoriteFlavours []string
}

func main() {
	person1 := person{
		first:            "Albert",
		last:             "Kolozsvari",
		favoriteFlavours: []string{"lemon", "strawberry", "vanilla"},
	}

	person2 := person{
		first:            "Joe",
		last:             "Doe",
		favoriteFlavours: []string{"chocolate"},
	}

	people := []person{
		person1, person2,
	}

	mappedPeople := make(map[string]person)

	for _, v := range people {
		mappedPeople[v.last] = v
	}

	for a, b := range mappedPeople {
		fmt.Println(a, "has properties:", b.first, b.last, b.favoriteFlavours)
	}
}
