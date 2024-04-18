package fly

type FlyWithWings struct{}

func (f *FlyWithWings) Fly() {
	println("I'm flying with my wings")
}

func NewFlyWithWings() *FlyWithWings {
	return &FlyWithWings{}
}
