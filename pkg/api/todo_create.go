package api

import (
	"encoding/json"
	"github.com/guilhermesteves/go-todo-api/internal/pkg/api/apimodel"
	"github.com/guilhermesteves/go-todo-api/internal/pkg/core/model"
	"log"

	"github.com/guilhermesteves/aclow"
	"github.com/guilhermesteves/go-todo-api/internal/pkg/core/message"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func createToDo(app *aclow.App, router *routing.Router) {
	handler := func(ctx *routing.Context) error {
		body := apimodel.CreateBody{}
		err := json.Unmarshal([]byte(ctx.PostBody()), &body)

		todo := &model.ToDo{ Name: body.Name }
		requestBody := message.CreatingToDoResquest{
			Todo: todo,
		}
		reply, err := app.Call("core@creating_todo", aclow.Message{Body: requestBody})

		ctx.SetContentType("application/json")

		if err != nil {
			log.Println(err.Error())
			responseBody, err := json.Marshal(map[string]string{"error": err.Error()})
			ctx.SetBody(responseBody)
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return err
		}

		replyBody := reply.Body.(message.CreatingToDoResponse).Todo
		responseBody, err := json.Marshal(apimodel.CreateResponse{
			ID: replyBody.Id,
		})

		ctx.SetBody(responseBody)
		ctx.SetStatusCode(fasthttp.StatusOK)
		return err
	}

	router.Post("/todos", handler)
}
