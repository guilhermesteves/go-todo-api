package flow

import (
	"log"
	"time"

	"github.com/guilhermesteves/aclow"
)

type EnsuringDatabaseStructure struct {
	app *aclow.App
}

func (n *EnsuringDatabaseStructure) Address() []string { return []string{"ensuring_database_structure"} }

func (n *EnsuringDatabaseStructure) Start(app *aclow.App) {
	n.app = app
	go func() {
		time.Sleep(time.Second * 5)
		app.Publish("data@ensuring_database_structure", aclow.Message{})
	}()
}

func (n *EnsuringDatabaseStructure) Execute(msg aclow.Message, call aclow.Caller) (aclow.Message, error) {
	log.Println("Ensuring DB structure...")

	call("data@create_todo_indexes", aclow.Message{})

	publish := n.app.Publish

	log.Println("Ensuring DB validations...")
	publish("data@create_todo_validation_schemas", aclow.Message{})
	log.Println("DB structure is ensured!")

	return aclow.Message{}, nil
}
