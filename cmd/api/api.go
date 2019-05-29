package main

import (
	"context"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/guilhermesteves/aclow"
	"github.com/guilhermesteves/go-todo-api/internal/pkg/util"
	"github.com/guilhermesteves/go-todo-api/pkg/api"
	"github.com/guilhermesteves/go-todo-api/pkg/core"
	"github.com/guilhermesteves/go-todo-api/pkg/data"
	"github.com/guilhermesteves/go-todo-api/pkg/data/config"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	godotenv.Load()

	startOpt := aclow.StartOptions{
		Local: true,
		Debug: false,
	}

	var app = &aclow.App{}

	app.Start(startOpt)

	router := routing.New()
	app.Resources["router"] = router
	connectOnMongo(app)

	app.RegisterModule("data", data.Nodes())
	app.RegisterModule("core", core.Nodes())

	api.RegisterRoutes(app)

	time.Sleep(time.Second * 2)

	httpServer(app, router)

	app.Wait()

}

func httpServer(app *aclow.App, router *routing.Router) {
	fasthttp.ListenAndServe("0.0.0.0:8080", router.HandleRequest)
}

var mongoTries = 0

func connectOnMongo(app *aclow.App) {
	opt := &options.ClientOptions{}
	client, err := mongo.NewClient(opt.ApplyURI(config.MongoDbDsn()))
	util.Check(err)

	err = client.Connect(context.TODO())

	if err != nil {
		time.Sleep(time.Second * 1)
		log.Println(err.Error())
		mongoTries++
		if mongoTries > 10 {
			log.Panic(err)
		}
		log.Println("Trying again...")
		connectOnMongo(app)
	} else {
		app.Resources["mongo"] = client
	}
}
