package duck

import (
	"example.com/strategy-pattern/fly"
	"example.com/strategy-pattern/quack"
)

type ModelDuck struct {
	Duck
}

func NewModelDuck() *ModelDuck {
	return &ModelDuck{
		Duck: Duck{
			flyBehavior:   fly.NewFlyNoWay(),
			quackBehavior: quack.NewQuack(),
		},
	}
}

func (d *ModelDuck) Display() {
	println("I'm a model duck")
}
