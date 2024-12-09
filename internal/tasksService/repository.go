package tasksService

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	GetAllTasks() ([]Task, error)
	CreateTask(task Task) (Task, error)
	UpdateTaskById(id uint, updatedTask Task) (Task, error)
	DeleteTaskById(id uint) error
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
func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}
func (r *taskRepository) UpdateTaskById(id uint, updatedTask Task) (Task, error) {
	findByID := r.db.First(&updatedTask, id)
	if findByID.Error != nil {
		return updatedTask, findByID.Error
	}
	result := r.db.Model(&Task{}).Where("id = ?", id).Update("task", updatedTask.Task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return updatedTask, nil
}
func (r *taskRepository) DeleteTaskById(id uint) error {
	var existingTask Task
	if err := r.db.First(&existingTask, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&Task{}, id).Error
}
