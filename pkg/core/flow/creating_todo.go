package flow

import "github.com/guilhermesteves/aclow"

type CreatingTodo struct {
	app  *aclow.App
}

func (n *CreatingTodo) Address() []string { return []string{"creating_todo"} }

func (n *CreatingTodo) Start(app *aclow.App) {
	n.app = app
}

func (n *CreatingTodo) Execute(msg aclow.Message, call aclow.Caller) (aclow.Message, error) {


	return aclow.Message{}, nil
}