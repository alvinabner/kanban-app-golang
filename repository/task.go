package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	task := []entity.Task{}
	err := r.db.WithContext(ctx).Where("user_id = ?", id).Find(&task).Error
	if err != nil {
		return []entity.Task{}, err
	}
	return task, nil // TODO: replace this
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	err = r.db.WithContext(ctx).Create(&task).Error
	if err != nil {
		return 0, err
	}
	return task.ID, nil // TODO: replace this
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	task := entity.Task{}
	err := r.db.WithContext(ctx).Where("id = ?", id).Limit(1).Find(&task).Error
	if err != nil {
		return entity.Task{}, err
	}
	return task, nil // TODO: replace this
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	task := []entity.Task{}
	err := r.db.WithContext(ctx).Where("category_id = ?", catId).Find(&task).Error
	if err != nil {
		return []entity.Task{}, err
	}
	return task, nil // TODO: replace this
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	err := r.db.WithContext(ctx).Updates(&task).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	category := entity.Task{}
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}
