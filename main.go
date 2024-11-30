package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
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
			Detail: "Could not add the message",
		})
	}
	if err := DB.Create(&task).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Could not create the message",
		})
	}
	return c.JSON(http.StatusOK, &task)
}
func PatchHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Invalid ID",
		})
	}
	var updatedTask Message
	if err := c.Bind(&updatedTask); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Could not update the message",
		})
	}

	if err := DB.Model(&Message{}).Where("id = ?", id).Update("task", updatedTask.Task).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Could not update the message",
		})
	}
	return c.JSON(http.StatusOK, Response{
		Status: "Success",
		Detail: "Message was successfully updated",
	})
}
func DeleteHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Invalid ID",
		})
	}
	if err := DB.Delete(&Message{}, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Could not delete the message",
		})
	}
	return c.JSON(http.StatusOK, Response{
		Status: "Success",
		Detail: "Message was successfully deleted",
	})
}

func main() {
	InitDB()

	if err := DB.AutoMigrate(&Message{}); err != nil {
		log.Fatalf("Не удалось мигрировать данные: %v", err)
	}

	c := echo.New()

	c.GET("/api/messages", GetHandler)
	c.POST("/api/messages", PostHandler)
	c.PATCH("/api/messages/:id", PatchHandler)
	c.DELETE("api/messages/:id", DeleteHandler)

	if err := c.Start(":8080"); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
