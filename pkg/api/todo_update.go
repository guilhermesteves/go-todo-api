package api

import (
	"encoding/json"
	"log"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/guilhermesteves/aclow"
	"github.com/guilhermesteves/go-todo-api/internal/pkg/core/message"
	"github.com/valyala/fasthttp"
)

func updateToDo(app *aclow.App, router *routing.Router) {
	handler := func(ctx *routing.Context) error {
		return err
	}

	router.Get("/todos", handler)
}
