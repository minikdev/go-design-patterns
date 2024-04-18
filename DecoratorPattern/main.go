package main

import (
	"fmt"

	"example.com/decorator-pattern/beverage"
)

func main() {
	b1 := beverage.NewEspresso()
	fmt.Printf("%s $%.2f\n", b1.GetDescription(), b1.Cost())
	mochaAdded := beverage.NewMocha(b1)
	fmt.Printf("%s $%.2f\n", mochaAdded.GetDescription(), mochaAdded.Cost())
}
