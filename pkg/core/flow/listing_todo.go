package flow

import "github.com/guilhermesteves/aclow"

type ListingTodo struct {
	app  *aclow.App
}

func (n *ListingTodo) Address() []string { return []string{"listing_todo"} }

func (n *ListingTodo) Start(app *aclow.App) {
	n.app = app
}

func (n *ListingTodo) Execute(msg aclow.Message, call aclow.Caller) (aclow.Message, error) {
	return aclow.Message{}, nil
}
