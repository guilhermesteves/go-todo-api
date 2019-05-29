package core

import (
	"github.com/guilhermesteves/aclow"
	"github.com/guilhermesteves/go-todo-api/pkg/core/flow"
	"github.com/guilhermesteves/go-todo-api/pkg/core/action"
)

func Nodes() []aclow.Node {
	return []aclow.Node{
		&flow.ListingTodo{},
		&flow.CreatingTodo{},
		&flow.LoadingTodo{},
		&flow.UpdatingTodo{},
		&flow.DeletingTodo{},

		&action.ListTodo{},
		&action.CreateTodo{},
		&action.LoadTodo{},
		&action.UpdateTodo{},
		&action.DeleteTodo{},
	}
}
