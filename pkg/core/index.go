package core

import (
	"github.com/guilhermesteves/aclow"
	"github.com/guilhermesteves/go-todo-api/pkg/core/flow"
)

func Nodes() []aclow.Node {
	return []aclow.Node{
		&flow.ListingTodo{},
		&flow.CreatingTodo{},
		&flow.LoadingTodo{},
		&flow.UpdatingTodo{},
		&flow.DeletingTodo{},
	}
}
