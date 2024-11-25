package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type response struct {
	Head string `json:"head"`
	Body string `json:"body"`
}
type requestBody struct {
	Message string `json:"message"`
}

var task string

func PostHandler(ctx echo.Context) error {
	var request requestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, response{
			Head: "Error",
			Body: "Message not added",
		})
	}
	task = request.Message
	return ctx.JSON(http.StatusOK, response{
		Head: "Success",
		Body: "Message was successfully added",
	})
}

func HelloHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello, "+task)
}
func main() {
	ctx := echo.New()
	ctx.GET("/api/hello", HelloHandler)
	ctx.POST("/api/hello", PostHandler)
	ctx.Start(":8080")
}
