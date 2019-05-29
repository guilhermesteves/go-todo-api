package action

import "github.com/guilhermesteves/aclow"

type DeleteTodo struct {
	app  *aclow.App
}

func (n *DeleteTodo) Address() []string { return []string{"delete_todo"} }

func (n *DeleteTodo) Start(app *aclow.App) {
	n.app = app
}

func (n *DeleteTodo) Execute(msg aclow.Message, call aclow.Caller) (aclow.Message, error) {
    return aclow.Message{}, nil
}