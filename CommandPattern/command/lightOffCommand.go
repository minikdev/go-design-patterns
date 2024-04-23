package command

import (
	"example.com/command-pattern/entities"
)

type LightOffCommand struct {
	Light entities.Light
}

func (l *LightOffCommand) Execute() {
	l.Light.Off()
}

func (l *LightOffCommand) GetName() string {
	return "LightOffCommand"
}

func (l *LightOffCommand) Undo() {
	l.Light.On()
}
