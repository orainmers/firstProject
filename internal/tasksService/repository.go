package tasksService

import (
	"firstProject/internal/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	GetAllTasks() ([]models.Task, error)
	CreateTask(task models.Task) (models.Task, error)
	UpdateTaskById(id uint, updatedTask models.Task) (models.Task, error)
	DeleteTaskById(id uint) error
	GetTasksByUserID(userID uint) ([]models.Task, error)
}
type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}
func (r *taskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil

}
func (r *taskRepository) CreateTask(task models.Task) (models.Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return models.Task{}, err
	}
	return task, nil
}
func (r *taskRepository) UpdateTaskById(id uint, updatedTask models.Task) (models.Task, error) {
	findByID := r.db.First(&models.Task{}, id)
	if findByID.Error != nil {
		return updatedTask, findByID.Error
	}
	updatedTask.ID = id
	result := r.db.Model(&updatedTask).Update("task", updatedTask.Task)
	if result.Error != nil {
		return models.Task{}, result.Error
	}
	return updatedTask, nil
}
func (r *taskRepository) DeleteTaskById(id uint) error {
	var existingTask models.Task
	if err := r.db.First(&existingTask, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&models.Task{}, id).Error
}
func (r *taskRepository) GetTasksByUserID(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
