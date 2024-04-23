package entities

import "fmt"

type Light struct {
	Name string
}

func (l *Light) On() {
	fmt.Println(l.Name, " Light is on")
}
func (l *Light) Off() {
	fmt.Println(l.Name, " Light is off")
}
