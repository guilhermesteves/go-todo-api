package api

import (
	"encoding/json"
	"github.com/guilhermesteves/go-todo-api/pkg/api/transformer"
	"log"

	"github.com/qiangxue/fasthttp-routing"
	"github.com/guilhermesteves/aclow"
	"github.com/guilhermesteves/go-todo-api/internal/pkg/core/message"
	"github.com/valyala/fasthttp"
)

func updateToDo(app *aclow.App, router *routing.Router) {
	handler := func(ctx *routing.Context) error {
		reply, err := app.Call("core@updating_todo", aclow.Message{}) // TODO: Implementar

		ctx.SetContentType("application/json")

		if err != nil {
			log.Println(err.Error())
			responseBody, err := json.Marshal(map[string]string{"error": err.Error()})
			ctx.SetBody(responseBody)
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return err
		}

		replyBody := reply.Body.(message.UpdatingToDoResponse).Todo
		responseBody, err := json.Marshal(transformer.ToDoToAPIToDo(replyBody))

		ctx.SetBody(responseBody)
		ctx.SetStatusCode(fasthttp.StatusOK)
		return err
	}

	router.Put("/todos/{id}", handler)
}
