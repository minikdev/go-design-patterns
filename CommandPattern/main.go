package main

import (
	"example.com/command-pattern/command"
	"example.com/command-pattern/entities"
	"example.com/command-pattern/remoteControl"
)

func main() {
	remote := remoteControl.SimpleRemoteControl{}
	lightOn := command.LightOnCommand{Light: entities.Light{Name: "Living Room"}}
	lightOff := command.LightOffCommand{Light: entities.Light{Name: "Living Room"}}
	kitchenLightOn := command.LightOnCommand{Light: entities.Light{Name: "Kitchen"}}
	kitchenLightOff := command.LightOffCommand{Light: entities.Light{Name: "Kitchen"}}
	remote.SetCommand(&lightOn)
	remote.ButtonWasPressed()

	remoteControl := remoteControl.NewRemoteControl()
	remoteControl.SetCommand(0, &lightOn, &lightOff)
	remoteControl.SetCommand(1, &kitchenLightOn, &kitchenLightOff)
	remoteControl.Display()

	remoteControl.OnButtonWasPushed(0)
	remoteControl.OffButtonWasPushed(0)
	remoteControl.UndoButtonWasPushed()
	remoteControl.OnButtonWasPushed(1)
	remoteControl.OffButtonWasPushed(1)
	remoteControl.OnButtonWasPushed(2)
	remoteControl.OffButtonWasPushed(2)
}
