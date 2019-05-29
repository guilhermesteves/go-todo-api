package flow

import "github.com/guilhermesteves/aclow"

type DeletingTodo struct {
	app  *aclow.App
}

func (n *DeletingTodo) Address() []string { return []string{"deleting_todo"} }

func (n *DeletingTodo) Start(app *aclow.App) {
	n.app = app
}

func (n *DeletingTodo) Execute(msg aclow.Message, call aclow.Caller) (aclow.Message, error) {


	return aclow.Message{}, nil
}