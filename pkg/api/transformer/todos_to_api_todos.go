package transformer

import (
	"github.com/guilhermesteves/go-todo-api/internal/pkg/api/apimodel"
	"github.com/guilhermesteves/go-todo-api/internal/pkg/core/model"
)

func ToDosToApiToDos(ToDos []*model.ToDo) []*apimodel.ToDo{
	var ToDosResponse []*apimodel.ToDo

	for _, td := range ToDos{
		t := ToDoToAPIToDo(td)
		ToDosResponse = append(ToDosResponse, t)
	}

	return ToDosResponse
}