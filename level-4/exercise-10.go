package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`: {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`: {`Being evil`, `Ice cream`, `Sunsets`},
	}

	x[`albert`] = []string{`martial arts`, `challenges`, `friends`}

	delete(x, `no_dr`)

	for a, b := range x {
		fmt.Println(a, "likes:")
		for c, d := range b {
			fmt.Println(c, d)
		}
	}
}
