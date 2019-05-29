package api

import (
	"bytes"
	"fmt"
	"log"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/guilhermesteves/aclow"
	"github.com/valyala/fasthttp"
)

func RegisterRoutes(app *aclow.App) {
	router := app.Resources["router"].(*routing.Router)

	router.Use(logHandler(), panicHandler(), corsHandler())

	listToDo(app, router)
	createToDo(app, router)
	loadToDo(app, router)
	updateToDo(app, router)
	deleteToDo(app, router)

}

func corsHandler() routing.Handler {
	return func(c *routing.Context) (err error) {
		c.Response.Header.Set("Access-Control-Allow-Origin", "*")
		c.Response.Header.Set("Access-Control-Allow-Headers", "Authorization")
		if bytes.Equal(c.Method(), []byte("OPTIONS")) {
			c.SetStatusCode(fasthttp.StatusOK)
			return nil
		}
		r := c.Next()
		return r
	}
}

func logHandler() routing.Handler {
	return func(c *routing.Context) (err error) {
		log.Println("Request: ", string(c.Method())+" "+string(c.Request.RequestURI()))
		r := c.Next()
		log.Println("Response Status: ", c.Response.StatusCode())
		return r
	}
}

func panicHandler() routing.Handler {
	return func(c *routing.Context) (err error) {
		defer func() {
			if e := recover(); e != nil {
				err := e.(error)
				c.Response.Header.SetStatusCode(fasthttp.StatusInternalServerError)
				eStr := fmt.Sprintf("%v", err)
				c.SetBody([]byte(eStr))
			}
		}()

		return c.Next()
	}
}
