package action

import "github.com/guilhermesteves/aclow"

type LoadTodo struct {
	app  *aclow.App
}

func (n *LoadTodo) Address() []string { return []string{"load_todo"} }

func (n *LoadTodo) Start(app *aclow.App) {
	n.app = app
}

func (n *LoadTodo) Execute(msg aclow.Message, call aclow.Caller) (aclow.Message, error) {
	return aclow.Message{}, nil
}
