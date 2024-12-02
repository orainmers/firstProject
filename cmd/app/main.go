package main

import (
	"firstProject/internal/database"
	"firstProject/internal/handlers"
	"firstProject/internal/taskService"
	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	c := echo.New()

	c.GET("/api/messages", handler.GetTasksHandler)
	c.POST("/api/messages", handler.PostTaskHandler)
	c.PATCH("/api/messages/:id", handler.UpdateTaskHandler)
	c.DELETE("/api/messages/:id", handler.DeleteTaskHandler)

	c.Start(":8080")
}
