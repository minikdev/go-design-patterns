package command

type ICommand interface {
	Execute()
	GetName() string
	Undo()
}
