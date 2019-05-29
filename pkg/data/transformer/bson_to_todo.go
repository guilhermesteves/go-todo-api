package transformer

import (
	"github.com/guilhermesteves/go-todo-api/internal/pkg/core/model"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BsonToToDo(d primitive.M) *model.ToDo {

	todo := &model.ToDo{
		Id:         cast.ToString(d["_id"]),
		Name:       cast.ToString(d["name"]),
		CreatedAt:  toFormattedTime(d["createdAt"]),
		UpdatedAt:	toFormattedTime(d["updatedAt"]),
	}

	return todo
}
