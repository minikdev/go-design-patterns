package main

type ITurkeyAdapter struct {
	turkey ITurkey
}

func (t *ITurkeyAdapter) Quack() {
	t.turkey.Gobble()
}

func (t *ITurkeyAdapter) Fly() {
	for i := 0; i < 5; i++ {
		t.turkey.Fly()
	}
}
