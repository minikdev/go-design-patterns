package fly

type FlyRocketPowered struct{}

func (f *FlyRocketPowered) Fly() {
	println("I'm flying with a rocket")
}
func NewFlyRocketPowered() *FlyRocketPowered {
	return &FlyRocketPowered{}
}
