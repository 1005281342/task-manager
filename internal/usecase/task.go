package usecase

import (
	"context"

	"github.com/1005281342/task-manager/internal/entity"
)

type TaskUC struct {
	repo TaskRepo
}

func NewTaskUC(repo TaskRepo) *TaskUC {
	return &TaskUC{repo: repo}
}

func (c *TaskUC) Index(ctx context.Context) (tasks []entity.Task, err error) {
	return c.repo.List(ctx)
}

func (c *TaskUC) Create(ctx context.Context, taskType string, cronSpec string, payload string) (*entity.Task, error) {
	var task = &entity.Task{
		CronSpec: cronSpec,
		Payload:  payload,
		Type:     taskType,
	}
	return c.repo.Create(ctx, task)
}

func (c *TaskUC) Show(ctx context.Context, id uint) (task *entity.Task, err error) {
	return c.repo.Get(ctx, id)
}

func (c *TaskUC) Edit(ctx context.Context, id uint) (task *entity.Task, err error) {
	return c.repo.Get(ctx, id)
}

func (c *TaskUC) Update(ctx context.Context, id uint, cronSpec string, payload string) error {
	return c.repo.Update(ctx, id, map[string]interface{}{
		"cron_spec": cronSpec,
		"payload":   payload,
	})
}

func (c *TaskUC) Delete(ctx context.Context, id uint) error {
	return c.repo.Delete(ctx, id)
}
