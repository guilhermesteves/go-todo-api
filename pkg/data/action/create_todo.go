package action

import (
	"context"
	"github.com/google/uuid"
	"github.com/guilhermesteves/aclow"
	"github.com/guilhermesteves/go-todo-api/internal/pkg/core/model"
	"github.com/guilhermesteves/go-todo-api/internal/pkg/util"
	"github.com/guilhermesteves/go-todo-api/pkg/data/config"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type CreateTodo struct {
	app  *aclow.App
}

func (n *CreateTodo) Address() []string { return []string{"create_todo"} }

func (n *CreateTodo) Start(app *aclow.App) {
	n.app = app
}

func (n *CreateTodo) Execute(msg aclow.Message, call aclow.Caller) (aclow.Message, error) {
	client := n.app.Resources["mongo"].(*mongo.Client)
	db := client.Database(config.MongoDbDatabase())

	name := msg.Body.(string)

	todo := model.ToDo{}
	todo.Id = uuid.New().String()
	todo.Name = name
	todo.CreatedAt = time.Now().Format(util.DefaultTimeMask)
	todo.UpdatedAt = time.Now().Format(util.DefaultTimeMask)

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	_, err := db.Collection("col_todo").InsertOne(ctx, todo)

	if err != nil {
		return aclow.Message{}, err
	}

	return aclow.Message{Body: todo}, nil
}