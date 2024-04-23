package main

import "fmt"

type ITurkey interface {
	Gobble()
	Fly()
}

type WildTurkey struct{}

func (w *WildTurkey) Gobble() {
	fmt.Println("Gobble gobble")
}

func (w *WildTurkey) Fly() {
	fmt.Println("I'm flying a short distance")
}
