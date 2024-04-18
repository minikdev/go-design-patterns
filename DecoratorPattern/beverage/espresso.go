package beverage

type Espresso struct {
	description string
}

func NewEspresso() *Espresso {
	return &Espresso{description: "Espresso"}
}
func (e *Espresso) Cost() float64 {
	return 1.99
}
func (e *Espresso) GetDescription() string {
	return e.description
}
