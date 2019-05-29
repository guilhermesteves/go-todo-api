package data

import (
	"github.com/guilhermesteves/aclow"
	"github.com/guilhermesteves/go-todo-api/pkg/data/action"
	"github.com/guilhermesteves/go-todo-api/pkg/data/flow"
)

func Nodes() []aclow.Node {
	return []aclow.Node{
		&action.CreateTodoIndexes{},
		&action.CreateTodoValidationSchemas{},

		&action.ListTodo{},
		&action.CreateTodo{},
		&action.LoadTodo{},
		&action.UpdateTodo{},
		&action.DeleteTodo{},

		&flow.EnsuringDatabaseStructure{},
	}

}
