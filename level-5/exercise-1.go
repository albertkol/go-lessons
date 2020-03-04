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

	fmt.Println(person1.first, person1.last, "likes:")
	for _, b := range person1.favoriteFlavours {
		fmt.Println(b)
	}

	fmt.Println(person2.first, person2.last, "likes:")
	for _, b := range person2.favoriteFlavours {
		fmt.Println(b)
	}
}
