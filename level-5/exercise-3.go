package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	car1 := truck{
		vehicle{doors: 6, color: "red"},
		false,
	}

	car2 := sedan{
		vehicle{doors:4, color:"yellow"},
		true,
	}

	fmt.Println(car1, car2)
	fmt.Println(car1.doors, car2.luxury)
}
