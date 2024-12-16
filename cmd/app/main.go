package main

import (
	"firstProject/internal/database"
	"firstProject/internal/handlers"
	"firstProject/internal/models"
	"firstProject/internal/tasksService"
	"firstProject/internal/userService"
	"firstProject/internal/web/tasks"
	"firstProject/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&models.Task{}); err != nil {
		log.Fatalf("Faild to migrate Task data: %v", err)
	}
	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Faild to migrate User data: %v", err)
	}

	repoTasks := tasksService.NewTaskRepository(database.DB)
	serviceTasks := tasksService.NewTaskService(repoTasks)
	handlerTasks := handlers.NewTaskHandler(serviceTasks)

	repoUsers := userService.NewUserRepository(database.DB)
	serviceUsers := userService.NewUserService(repoUsers)
	handlerUsers := handlers.NewUserHandler(serviceUsers)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	taskHandler := tasks.NewStrictHandler(handlerTasks, nil)
	tasks.RegisterHandlers(e, taskHandler)
	userHandler := users.NewStrictHandler(handlerUsers, nil)
	users.RegisterHandlers(e, userHandler)

	if err := e.Start(":8081"); err != nil {
		log.Fatalf("Failed to start the server with err: %v", err)
	}
}
