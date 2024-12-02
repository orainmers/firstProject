package taskService

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	GetAllTasks() ([]Task, error)
	CreateTask(task *Task) error
	UpdateTaskById(id uint, updatedTask *Task) error
	DeleteTaskById(id uint) error
}
type Response struct {
	Status string `json:"status"`
	Detail string `json:"detail"`
}
type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}
func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil

}
func (r *taskRepository) CreateTask(task *Task) error {
	return r.db.Create(&task).Error
}
func (r *taskRepository) UpdateTaskById(id uint, updatedTask *Task) error {
	return r.db.Model(&Task{}).Where("id = ?", id).Update("task", updatedTask.Task).Error
}
func (r *taskRepository) DeleteTaskById(id uint) error {
	return r.db.Delete(&Task{}, id).Error
}
