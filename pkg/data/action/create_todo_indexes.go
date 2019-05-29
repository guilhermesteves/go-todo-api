package action

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/guilhermesteves/aclow"
	"github.com/guilhermesteves/go-todo-api/pkg/data/config"
)

type CreateTodoIndexes struct {
	app *aclow.App
}

func (n *CreateTodoIndexes) Address() []string { return []string{"create_todo_indexes"} }

func (n *CreateTodoIndexes) Start(app *aclow.App) {
	n.app = app
}

func (n *CreateTodoIndexes) Execute(msg aclow.Message, call aclow.Caller) (aclow.Message, error) {
	client := n.app.Resources["mongo"].(*mongo.Client)
	db := client.Database(config.MongoDbDatabase())

	db.Collection("col_event").Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		mongo.IndexModel{Keys: bson.M{"name": 1}},
	}, &options.CreateIndexesOptions{})

	return aclow.Message{}, nil
}
