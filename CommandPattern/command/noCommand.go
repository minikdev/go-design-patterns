package command

type NoCommand struct{}

func (n *NoCommand) Execute() {
	// Do nothing
}

func (n *NoCommand) GetName() string {
	return "NoCommand"
}

func (n *NoCommand) Undo() {
	// Do nothing
}
