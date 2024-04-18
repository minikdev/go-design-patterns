package quack

type Quack struct {
}

func (q *Quack) Quack() {
	println("Quack quack")
}

func NewQuack() *Quack {
	return &Quack{}
}
