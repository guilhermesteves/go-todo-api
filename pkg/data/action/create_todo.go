package action

import "github.com/guilhermesteves/aclow"

type CreateTodo struct {
	app  *aclow.App
}

func (n *CreateTodo) Address() []string { return []string{"create_todo"} }

func (n *CreateTodo) Start(app *aclow.App) {
	n.app = app
}

func (n *CreateTodo) Execute(msg aclow.Message, call aclow.Caller) (aclow.Message, error) {
    return aclow.Message{}, nil
}