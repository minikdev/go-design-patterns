package main

import (
	"fmt"

	"example.com/factory-method-pattern/gun"
)

func main() {
	ak47, _ := gun.GetGun("ak47")
	baretta, _ := gun.GetGun("baretta")
	fmt.Println("1st Gun:", ak47.GetName(), ak47.GetPower())
	fmt.Println("2nd Gun:", baretta.GetName(), baretta.GetPower())

}
