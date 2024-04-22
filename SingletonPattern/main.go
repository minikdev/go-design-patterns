package main

import (
	"example.com/singleton-pattern/chocolateboiler"
)

func main() {
	cb := chocolateboiler.GetInstance()
	cb.Fill()
	cb.Boil()
	cb.Drain()
	cb = chocolateboiler.GetInstance()
	cb.Boil()
	cb.Drain()
}
