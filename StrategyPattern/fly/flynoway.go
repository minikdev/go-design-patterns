package fly

type FlyNoWay struct{}

func (f *FlyNoWay) Fly() {
	println("I can't fly")
}

func NewFlyNoWay() *FlyNoWay {
	return &FlyNoWay{}
}
