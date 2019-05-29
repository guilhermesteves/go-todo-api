package api

import (
	"encoding/json"
	"log"

	"github.com/qiangxue/fasthttp-routing"
	"github.com/guilhermesteves/aclow"
	"github.com/guilhermesteves/go-todo-api/internal/pkg/core/message"
	"github.com/guilhermesteves/go-todo-api/pkg/api/transformer"
	"github.com/valyala/fasthttp"
)

func listToDo(app *aclow.App, router *routing.Router) {
	handler := func(ctx *routing.Context) error {
		reply, err := app.Call("core@listing_todos", aclow.Message{})

		ctx.SetContentType("application/json")

		if err != nil {
			log.Println(err.Error())
			responseBody, err := json.Marshal(map[string]string{"error": err.Error()})
			ctx.SetBody(responseBody)
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return err
		}

		replyBody := reply.Body.(message.ListingToDoResponse).Todos
		responseBody, err := json.Marshal(transformer.ToDosToApiToDos(replyBody))

		ctx.SetBody(responseBody)
		ctx.SetStatusCode(fasthttp.StatusOK)
		return err
	}

	router.Get("/todos", handler)
}
