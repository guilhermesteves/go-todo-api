package flow

import "github.com/guilhermesteves/aclow"

type UpdatingTodo struct {
	app  *aclow.App
}

func (n *UpdatingTodo) Address() []string { return []string{"updating_todo"} }

func (n *UpdatingTodo) Start(app *aclow.App) {
	n.app = app
}

func (n *UpdatingTodo) Execute(msg aclow.Message, call aclow.Caller) (aclow.Message, error) {
    return aclow.Message{}, nil
}