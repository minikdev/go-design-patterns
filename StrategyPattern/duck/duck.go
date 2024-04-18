package duck

import (
	"example.com/strategy-pattern/duckbehaviors"
)

type Duck struct {
	quackBehavior duckbehaviors.IQuackBehavior
	flyBehavior   duckbehaviors.IFlyBehavior
}

func (d *Duck) PerformQuack() {
	d.quackBehavior.Quack()
}
func (d *Duck) PerformFly() {
	d.flyBehavior.Fly()
}
func (d *Duck) Swim() {
	println("All ducks float, even decoys!")
}
func (d *Duck) Display() {
	println("I'm a duck!")
}
func (d *Duck) SetQuackBehavior(quackBehavior duckbehaviors.IQuackBehavior) {
	d.quackBehavior = quackBehavior
}
func (d *Duck) SetFlyBehavior(flyBehavior duckbehaviors.IFlyBehavior) {
	d.flyBehavior = flyBehavior
}
