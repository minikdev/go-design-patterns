package main

import (
	"example.com/strategy-pattern/duck"
	"example.com/strategy-pattern/fly"
)

func main() {
	mallardDuck := duck.NewMallardDuck()
	mallardDuck.Display()
	mallardDuck.PerformQuack()
	mallardDuck.PerformFly()

	modelDuck := duck.NewModelDuck()
	modelDuck.Display()
	modelDuck.PerformQuack()
	modelDuck.PerformFly()
	modelDuck.SetFlyBehavior(fly.NewFlyRocketPowered())
	modelDuck.PerformFly()
}
