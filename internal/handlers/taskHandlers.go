package handlers

import (
	"context"
	"firstProject/internal/models"
	"firstProject/internal/tasksService"
	"firstProject/internal/web/tasks"
	"gorm.io/gorm"
)

type TaskHandler struct {
	Service *tasksService.TaskService
}

func NewTaskHandler(service *tasksService.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

func (t *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := t.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}
	var response tasks.GetTasks200JSONResponse
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			IsDone: &tsk.IsDone,
			Task:   &tsk.Task,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}
	return response, nil
}

func (t *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := models.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
		UserID: *taskRequest.UserId,
	}
	createdTask, err := t.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Task:   &createdTask.Task,
		Id:     &createdTask.ID,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}
	return response, nil
}

func (t *TaskHandler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskID := uint(request.Id)
	taskRequest := request.Body
	taskToUpdate := models.Task{
		Model:  gorm.Model{ID: taskID},
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	updatedTask, err := t.Service.UpdateTaskById(taskID, taskToUpdate)
	if err != nil {
		return nil, err
	}
	updatedTask.ID = taskID
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		IsDone: &updatedTask.IsDone,
		Task:   &updatedTask.Task,
	}
	return response, nil
}

func (t *TaskHandler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id
	if err := t.Service.DeleteTaskById(uint(taskID)); err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204JSONResponse{}, nil
}
func (t *TaskHandler) GetTasksUserId(_ context.Context, request tasks.GetTasksUserIdRequestObject) (tasks.GetTasksUserIdResponseObject, error) {
	userID := request.UserId
	userTasks, err := t.Service.GetTasksByUserID(userID)
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasksUserId200JSONResponse{}
	for _, task := range userTasks {
		tsk := tasks.Task{
			Id:     &task.ID,
			Task:   &task.Task,
			IsDone: &task.IsDone,
			UserId: &task.UserID,
		}
		response = append(response, tsk)
	}
	return response, nil
}
func (t *TaskHandler) GetUsersUserIdTasks(_ context.Context, request tasks.GetUsersUserIdTasksRequestObject) (tasks.GetUsersUserIdTasksResponseObject, error) {
	userID := request.UserId
	userTasks, err := t.Service.GetTasksByUserID(userID)
	if err != nil {
		return nil, err
	}

	response := tasks.GetUsersUserIdTasks200JSONResponse{}

	for _, tsk := range userTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}
