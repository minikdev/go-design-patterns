package beverage

type Mocha struct {
	beverage Beverage
}

func NewMocha(b Beverage) *Mocha {
	return &Mocha{beverage: b}
}
func (m *Mocha) GetDescription() string {
	return m.beverage.GetDescription() + ", Mocha"
}
func (m *Mocha) Cost() float64 {
	return 0.20 + m.beverage.Cost()
}
