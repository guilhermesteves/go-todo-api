package api

import (
	"encoding/json"
	"log"

	"github.com/guilhermesteves/aclow"
	"github.com/guilhermesteves/go-todo-api/internal/pkg/core/message"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func deleteToDo(app *aclow.App, router *routing.Router) {
	handler := func(ctx *routing.Context) error {
		reply, err := app.Call("core@deleting_todo", aclow.Message{}) // TODO: Implementar

		ctx.SetContentType("application/json")

		if err != nil {
			log.Println(err.Error())
			responseBody, err := json.Marshal(map[string]string{"error": err.Error()})
			ctx.SetBody(responseBody)
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return err
		}

		replyBody := reply.Body.(message.DeletingToDoResponse).Success
		success := bool(replyBody)

		if success != true {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
		} else {
			ctx.SetStatusCode(fasthttp.StatusOK)
		}

		return err
	}

	router.Get("/todos", handler)
}
