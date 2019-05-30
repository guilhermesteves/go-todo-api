package action

import (
	"context"
	"github.com/guilhermesteves/aclow"
	"github.com/guilhermesteves/go-todo-api/pkg/data/config"
	"time"

	dbtransformer "github.com/guilhermesteves/go-todo-api/pkg/data/transformer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoadTodo struct {
	app  *aclow.App
}

func (n *LoadTodo) Address() []string { return []string{"load_todo"} }

func (n *LoadTodo) Start(app *aclow.App) {
	n.app = app
}

func (n *LoadTodo) Execute(msg aclow.Message, call aclow.Caller) (aclow.Message, error) {
	client := n.app.Resources["mongo"].(*mongo.Client)
	db := client.Database(config.MongoDbDatabase())

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	result := db.Collection("col_todo").FindOne(
		ctx,
		bson.M{ "_id": msg.Body.(string), },
		&options.FindOneOptions{ Projection: bson.M{},
	})

	if result.Err() != nil {
		return aclow.Message{}, result.Err()
	}

	var todo primitive.M
	err := result.Decode(&todo)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return aclow.Message{}, nil
		}

		return aclow.Message{}, err
	}

	return aclow.Message{Body: dbtransformer.BsonToToDo(todo)}, nil
}
