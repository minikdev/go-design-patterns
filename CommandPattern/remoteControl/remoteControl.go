package remoteControl

import (
	"errors"
	"fmt"

	"example.com/command-pattern/command"
)

type RemoteControl struct {
	onCommands  []command.ICommand
	offCommands []command.ICommand
	undoCommand command.ICommand
}

func NewRemoteControl() *RemoteControl {
	onCommands := make([]command.ICommand, 7)
	offCommands := make([]command.ICommand, 7)

	for i := 0; i < len(onCommands); i++ {
		onCommands[i] = &command.NoCommand{}
		offCommands[i] = &command.NoCommand{}
	}

	return &RemoteControl{
		onCommands:  onCommands,
		offCommands: offCommands,
		undoCommand: &command.NoCommand{},
	}
}

func (r *RemoteControl) SetCommand(slot int, onCommand command.ICommand, offCommand command.ICommand) error {
	if slot < 0 || slot >= len(r.onCommands) {
		return errors.New("invalid slot")
	}
	r.onCommands[slot] = onCommand
	r.offCommands[slot] = offCommand
	return nil
}

func (r *RemoteControl) OnButtonWasPushed(slot int) error {
	if slot < 0 || slot >= len(r.onCommands) {
		return errors.New("invalid slot")
	}
	r.onCommands[slot].Execute()
	r.undoCommand = r.onCommands[slot]
	return nil
}

func (r *RemoteControl) OffButtonWasPushed(slot int) error {
	if slot < 0 || slot >= len(r.offCommands) {
		return errors.New("invalid slot")
	}
	r.offCommands[slot].Execute()
	r.undoCommand = r.offCommands[slot]
	return nil
}

func (r *RemoteControl) Display() {
	fmt.Println("-------------- Remote Control --------------")
	for i := 0; i < len(r.onCommands); i++ {
		fmt.Printf("[slot %d] %s\t%s\n", i, r.onCommands[i].GetName(), r.offCommands[i].GetName())
	}
	fmt.Println("--------------------------------------------")
}

func (r *RemoteControl) UndoButtonWasPushed() {
	r.undoCommand.Undo()
}
