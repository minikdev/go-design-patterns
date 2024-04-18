package duck

import (
	"example.com/strategy-pattern/fly"
	"example.com/strategy-pattern/quack"
)

type MallardDuck struct {
	Duck
}

func NewMallardDuck() *MallardDuck {
	return &MallardDuck{
		Duck: Duck{
			quackBehavior: quack.NewQuack(),
			flyBehavior:   fly.NewFlyWithWings(),
		},
	}
}

func (d *MallardDuck) Display() {
	println("I'm a real Mallard duck")
}
