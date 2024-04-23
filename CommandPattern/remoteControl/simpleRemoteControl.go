package remoteControl

import "example.com/command-pattern/command"

type SimpleRemoteControl struct {
	slot command.ICommand
}

func (s *SimpleRemoteControl) SetCommand(command command.ICommand) {
	s.slot = command
}
func (s *SimpleRemoteControl) ButtonWasPressed() {
	s.slot.Execute()
}
