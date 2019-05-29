package flow

import (
	"fmt"
	"github.com/guilhermesteves/aclow"
	"github.com/guilhermesteves/go-todo-api/internal/pkg/core/message"
	"github.com/guilhermesteves/go-todo-api/internal/pkg/core/model"
)

type CreatingTodo struct {
	app  *aclow.App
}

func (n *CreatingTodo) Address() []string { return []string{"creating_todo"} }

func (n *CreatingTodo) Start(app *aclow.App) {
	n.app = app
}

func (n *CreatingTodo) Execute(msg aclow.Message, call aclow.Caller) (aclow.Message, error) {
	body := msg.Body.(message.CreatingToDoResquest)

	reply, _ := call("data@create_todo", aclow.Message{Body: body.Todo.Name})

	if reply.Body == nil {
		return aclow.Message{}, fmt.Errorf("ERROR")
	}

	todo := reply.Body.(model.ToDo)

	return aclow.Message{Body: message.CreatingToDoResponse{
		Todo: &todo,
	}}, nil
}