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

	body := msg.Body.(map[string]string)

	batchSize := int32(1)

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	cur, err := db.Collection("col_todo").Aggregate(ctx, bson.A{
		bson.M{"$match": bson.M{
			"_id":       body["personID"],
			"deals._id": body["dealID"],
		}},
		bson.M{"$unwind": "$deals"},
		bson.M{"$match": bson.M{
			"deals._id": body["dealID"],
		}},
		bson.M{"$replaceRoot": bson.M{
			"newRoot": "$deals",
		}},
		bson.M{"$limit": 1},
	}, &options.AggregateOptions{
		BatchSize: &batchSize,
	})

	if err != nil {
		return aclow.Message{}, err
	}

	var todo primitive.M
	defer cur.Close(ctx)

	if cur.Next(ctx) {
		cur.Decode(&todo)
	}

	if todo["_id"] != nil && todo["_id"] != "" {
		d := dbtransformer.BsonToToDo(todo)
		return aclow.Message{Body: *d}, nil
	} else {
		return aclow.Message{}, nil
	}
}
