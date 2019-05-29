package action

import "github.com/guilhermesteves/aclow"

type UpdateTodo struct {
	app  *aclow.App
}

func (n *UpdateTodo) Address() []string { return []string{"update_todo"} }

func (n *UpdateTodo) Start(app *aclow.App) {
	n.app = app
}

func (n *UpdateTodo) Execute(msg aclow.Message, call aclow.Caller) (aclow.Message, error) {
    return aclow.Message{}, nil
}