package command

import "example.com/command-pattern/entities"

type LightOnCommand struct {
	Light entities.Light
}

func (l *LightOnCommand) Execute() {
	l.Light.On()
}

func (l *LightOnCommand) GetName() string {
	return "LightOnCommand"
}

func (l *LightOnCommand) Undo() {
	l.Light.Off()
}
