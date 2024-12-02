package main

import (
	"firstProject/internal/database"
	"firstProject/internal/taskService"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Response struct {
	Status string `json:"status"`
	Detail string `json:"detail"`
}

func GetHandler(c echo.Context) error {
	var messages []taskService.Message
	if err := database.DB.Find(&messages).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Could not find the messages",
		})
	}
	return c.JSON(http.StatusOK, &messages)
}
func PostHandler(c echo.Context) error {
	var task taskService.Message
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Could not add the message",
		})
	}
	if err := database.DB.Create(&task).Error; err != nil {
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
	var updatedTask taskService.Message
	if err := c.Bind(&updatedTask); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Could not update the message",
		})
	}

	if err := database.DB.Model(&taskService.Message{}).Where("id = ?", id).Update("task", updatedTask.Task).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Could not update the message",
		})
	}
	return c.JSON(http.StatusOK, &updatedTask)
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
	if err := database.DB.Delete(&taskService.Message{}, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Could not delete the message",
		})
	}
	return c.NoContent(http.StatusNoContent)
}

func main() {
	database.InitDB()

	database.DB.AutoMigrate(&taskService.Message{})

	c := echo.New()

	c.GET("/api/messages", GetHandler)
	c.POST("/api/messages", PostHandler)
	c.PATCH("/api/messages/:id", PatchHandler)
	c.DELETE("/api/messages/:id", DeleteHandler)

	c.Start(":8080")
}
