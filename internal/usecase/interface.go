package usecase

import (
	"context"

	"github.com/hibiken/asynq"

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

type ListRepo interface {
	List(context.Context) ([]entity.Task, error)
}

type TaskRepo interface {
	ListRepo
	Create(context.Context, *entity.Task) (*entity.Task, error)
	Get(context.Context, uint) (*entity.Task, error)
	Update(context.Context, uint, map[string]interface{}) error
	Delete(context.Context, uint) error
}

type Scheduler interface {
	GetConfigs() ([]*asynq.PeriodicTaskConfig, error)
}

type SchedulerRepo interface {
	ListRepo
}
