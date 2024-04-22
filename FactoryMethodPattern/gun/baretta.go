package gun

type Baretta struct {
	Gun
}

func NewBaretta() IGun {
	return &Baretta{
		Gun: Gun{
			name:  "Baretta gun",
			power: 3,
		},
	}
}
