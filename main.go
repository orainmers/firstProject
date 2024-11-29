package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Detail string `json:"detail"`
}

func GetHandler(c echo.Context) error {
	var messages []Message
	if err := DB.Find(&messages).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Could not find the messages",
		})
	}
	return c.JSON(http.StatusOK, &messages)
}
func PostHandler(c echo.Context) error {
	var task Message
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Message not added",
		})
	}
	if err := DB.Create(&task).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Could not create the message",
		})
	}
	return c.JSON(http.StatusOK, Response{
		Status: "Success",
		Detail: "Message was successfully added",
	})
}

func main() {
	InitDB()

	if err := DB.AutoMigrate(&Message{}); err != nil {
		log.Fatalf("Не удалось мигрировать данные: %v", err)
	}

	c := echo.New()

	c.GET("/api/hello", GetHandler)
	c.POST("/api/hello", PostHandler)

	if err := c.Start(":8080"); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
