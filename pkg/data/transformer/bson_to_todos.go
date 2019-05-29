package transformer

import (
	"github.com/guilhermesteves/go-todo-api/internal/pkg/core/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BsonToToDos(src primitive.A) []*model.ToDo {
	todos := []*model.ToDo{}

	if src != nil {
		for _, d := range src {
			todos = append(todos, BsonToToDo(d.(primitive.M)))
		}
	}

	return todos
}