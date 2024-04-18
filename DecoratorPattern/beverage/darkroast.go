package beverage

type DarkRoast struct {
	description string
}

func NewDarkRoast() *DarkRoast {
	return &DarkRoast{description: "Dark Roast Coffee"}
}

func (d *DarkRoast) Cost() float64 {
	return 0.99
}
