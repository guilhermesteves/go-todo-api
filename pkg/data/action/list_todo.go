package action

import (
	"context"
	"github.com/guilhermesteves/aclow"
	"github.com/guilhermesteves/go-todo-api/pkg/data/config"
	dbtransformer "github.com/guilhermesteves/go-todo-api/pkg/data/transformer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ListTodo struct {
	app  *aclow.App
}

func (n *ListTodo) Address() []string { return []string{"list_todo"} }

func (n *ListTodo) Start(app *aclow.App) {
	n.app = app
}

func (n *ListTodo) Execute(msg aclow.Message, call aclow.Caller) (aclow.Message, error) {
	client := n.app.Resources["mongo"].(*mongo.Client)
	db := client.Database(config.MongoDbDatabase())

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)

	cur, err := db.Collection("col_todo").Find(ctx, bson.M{}, &options.FindOptions{})

	if err != nil {

	}

	var todos primitive.A
	cur.Decode(&todos)

	if len(todos) > 0 {
		return aclow.Message{Body: dbtransformer.BsonToToDos(todos)}, nil
	} else {
		return aclow.Message{}, nil
	}
}
