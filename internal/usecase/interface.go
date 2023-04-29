package usecase

import (
	"context"

	"github.com/1005281342/task-manager/internal/entity"
)

type Task interface {
	Index(ctx context.Context) (tasks []entity.Task, err error)
	Create(ctx context.Context, taskType string, cronSpec string, payload string) (task *entity.Task, err error)
	Show(ctx context.Context, id uint) (task *entity.Task, err error)
	Edit(ctx context.Context, id uint) (task *entity.Task, err error)
	Update(ctx context.Context, id uint, cronSpec string, payload string) error
	Delete(ctx context.Context, id uint) error
}

type TaskRepo interface {
	List(context.Context) ([]entity.Task, error)
	Create(context.Context, *entity.Task) (*entity.Task, error)
	Get(context.Context, uint) (*entity.Task, error)
	Update(context.Context, uint, map[string]interface{}) error
	Delete(context.Context, uint) error
}
