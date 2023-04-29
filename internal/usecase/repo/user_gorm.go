package repo

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/1005281342/task-manager/internal/entity"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (u *TaskRepo) List(ctx context.Context) ([]entity.Task, error) {
	var (
		tasks []entity.Task
		err   error
	)
	if err = u.db.Find(&tasks).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return tasks, nil
}

func (u *TaskRepo) Create(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	if task == nil {
		return nil, fmt.Errorf("task is nil")
	}

	if err := u.db.Model(&entity.Task{}).Create(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (u *TaskRepo) Get(ctx context.Context, id uint) (*entity.Task, error) {
	if id == 0 {
		return nil, fmt.Errorf("id is 0")
	}

	var task = &entity.Task{}
	if err := u.db.Model(&entity.Task{}).Where("id = ?", id).Find(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (u *TaskRepo) Update(ctx context.Context, id uint, task map[string]interface{}) error {
	return u.db.Model(&entity.Task{}).Where("id = ?", id).Updates(task).Error
}

func (u *TaskRepo) Delete(ctx context.Context, id uint) error {
	return u.db.Model(&entity.Task{}).Where("id = ?", id).Delete(&entity.Task{Model: gorm.Model{ID: id}}).Error
}
