package handlers

import (
	"firstProject/internal/taskService"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	Status string `json:"status"`
	Detail string `json:"detail"`
}

type TaskHandler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

func (h *TaskHandler) GetTasksHandler(c echo.Context) error {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Could not find the tasks",
		})
	}
	return c.JSON(http.StatusOK, &tasks)
}
func (h *TaskHandler) PostTaskHandler(c echo.Context) error {
	var newTask taskService.Task
	if err := c.Bind(&newTask); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Could not add the task",
		})
	}
	if err := h.Service.CreateTask(&newTask); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Could not create the task",
		})
	}
	return c.JSON(http.StatusOK, &newTask)
}
func (h *TaskHandler) UpdateTaskHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Invalid ID",
		})
	}

	var updatedTask taskService.Task
	if err := c.Bind(&updatedTask); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Could not bind the task",
		})
	}

	updatedTask.ID = uint(id)
	updatedTask.UpdatedAt = time.Now()
	
	if err := h.Service.UpdateTaskById(uint(id), &updatedTask); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Could not update the task",
		})
	}
	return c.JSON(http.StatusOK, &updatedTask)
}
func (h *TaskHandler) DeleteTaskHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Invalid ID",
		})
	}
	if err := h.Service.DeleteTaskById(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Detail: "Could not delete the task",
		})
	}
	return c.NoContent(http.StatusNoContent)
}
