package action

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/guilhermesteves/aclow"
	"github.com/guilhermesteves/go-todo-api/internal/pkg/data/schema"
	"github.com/guilhermesteves/go-todo-api/pkg/data/config"
)

type CreateTodoValidationSchemas struct {
	app *aclow.App
}

func (n *CreateTodoValidationSchemas) Address() []string {
	return []string{"create_todo_validation_schemas"}
}

func (n *CreateTodoValidationSchemas) Start(app *aclow.App) {
	n.app = app
}

func (n *CreateTodoValidationSchemas) Execute(msg aclow.Message, call aclow.Caller) (aclow.Message, error) {
	client := n.app.Resources["mongo"].(*mongo.Client)
	db := client.Database(config.MongoDbDatabase())

	jsonSchema := bson.M{}
	jsonFile, _ := schema.ApiDataSchemaTodoJsonBytes()

	err := json.Unmarshal(jsonFile, &jsonSchema)

	if err != nil {
		return aclow.Message{}, err
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	result := db.RunCommand(ctx, bson.D{
		{"collMod", "col_todo"},
		{"validator", bson.M{"$jsonSchema": jsonSchema}},
		{"validationLevel", "moderate"},
	}, &options.RunCmdOptions{})

	if result.Err() != nil {
		log.Println(n.Address(), result.Err())
	}

	return aclow.Message{}, nil
}
