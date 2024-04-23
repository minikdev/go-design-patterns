package main

func main() {
	mallardDuck := &MallardDuck{}
	wildTurkey := &WildTurkey{}
	turkeyAdapter := &ITurkeyAdapter{turkey: wildTurkey}

	mallardDuck.Quack()
	mallardDuck.Fly()
	turkeyAdapter.Quack()
	turkeyAdapter.Fly()
}
