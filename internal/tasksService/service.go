package tasksService

import "firstProject/internal/models"

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}
func (s *TaskService) CreateTask(task models.Task) (models.Task, error) {
	return s.repo.CreateTask(task)
}
func (s *TaskService) UpdateTaskById(id uint, updatedTask models.Task) (models.Task, error) {
	return s.repo.UpdateTaskById(id, updatedTask)
}
func (s *TaskService) DeleteTaskById(id uint) error {
	return s.repo.DeleteTaskById(id)
}
func (s *TaskService) GetTasksByUserID(userID uint) ([]models.Task, error) {
	return s.repo.GetTasksByUserID(userID)
}
