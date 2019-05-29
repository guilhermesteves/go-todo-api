package transformer

import (
	"github.com/guilhermesteves/go-todo-api/internal/pkg/api/apimodel"
	"github.com/guilhermesteves/go-todo-api/internal/pkg/core/model"
)

func ToDoToAPIToDo(ToDo *model.ToDo) *apimodel.ToDo{
	apiToDo := &apimodel.ToDo{
		ID: 		ToDo.Id,
		Name: 		ToDo.Name,
		CreatedAt: 	ToDo.CreatedAt,
		UpdatedAt:	ToDo.UpdatedAt,
	}

	return apiToDo
}