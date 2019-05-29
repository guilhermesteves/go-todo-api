package flow

import "github.com/guilhermesteves/aclow"

type LoadingTodo struct {
	app  *aclow.App
}

func (n *LoadingTodo) Address() []string { return []string{"loading_todo"} }

func (n *LoadingTodo) Start(app *aclow.App) {
	n.app = app
}

func (n *LoadingTodo) Execute(msg aclow.Message, call aclow.Caller) (aclow.Message, error) {
	return aclow.Message{}, nil
}