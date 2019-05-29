package action

import "github.com/guilhermesteves/aclow"

type ListTodo struct {
	app  *aclow.App
}

func (n *ListTodo) Address() []string { return []string{"list_todo"} }

func (n *ListTodo) Start(app *aclow.App) {
	n.app = app
}

func (n *ListTodo) Execute(msg aclow.Message, call aclow.Caller) (aclow.Message, error) {
	return aclow.Message{}, nil
}
