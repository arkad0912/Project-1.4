package taskService

import (
	"time"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTaskByID(id uint, task Task) (Task, error)
	DeleteTaskByID(id uint) (time.Time, error)    // Обновлено
	GetTasksByUserID(userID uint) ([]Task, error) // Добавьте этот метод
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTasksByUserID(userID uint) ([]Task, error) {
	var tasks []Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	var existingTask Task
	if err := r.db.First(&existingTask, id).Error; err != nil {
		return Task{}, err
	}

	if task.Task != "" {
		existingTask.Task = task.Task
	}
	if task.IsDone != existingTask.IsDone {
		existingTask.IsDone = task.IsDone
	}

	if err := r.db.Save(&existingTask).Error; err != nil {
		return Task{}, err
	}

	return existingTask, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) (time.Time, error) {
	// Получаем текущее время (дату удаления)
	deletedAt := time.Now()

	// Выполняем физическое удаление
	result := r.db.Model(&Task{}).Where("id = ?", id).Update("deleted_at", deletedAt)
	if result.Error != nil {
		return time.Time{}, result.Error
	}

	// Возвращаем дату удаления
	return deletedAt, nil
}
