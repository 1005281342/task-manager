package task

import (
	"context"
	"sync"

	"github.com/1005281342/task-manager/internal/config"
	"github.com/1005281342/task-manager/internal/entity"
	"github.com/1005281342/task-manager/internal/taskscheduler"
	"github.com/1005281342/task-manager/internal/usecase"
	"github.com/1005281342/task-manager/internal/usecase/repo"
	"github.com/1005281342/task-manager/pkg/db"
)

// Controller for tasks
type Controller struct {
	config.Config
	uc usecase.Task
}

var once sync.Once

func Load(cfg config.Config) (*Controller, error) {
	var t, err = db.New(cfg.Gorm.Driver, cfg.Gorm.Dsn)
	if err != nil {
		return nil, err
	}
	if err = t.AutoMigrate(&entity.Task{}); err != nil {
		return nil, err
	}

	var uc = usecase.NewTaskUC(repo.NewTaskRepo(t))

	once.Do(func() {
		go taskscheduler.Run(cfg)
	})

	return &Controller{uc: uc}, nil
}

// Index of tasks
// GET /task
func (c *Controller) Index(ctx context.Context) ([]entity.Task, error) {
	return c.uc.Index(ctx)
}

// New returns a view for creating a new task
// GET /task/new
func (c *Controller) New(ctx context.Context) {
}

// Create task
// POST /task
func (c *Controller) Create(ctx context.Context, taskType string, cronSpec string, payload string) (*entity.Task, error) {
	return c.uc.Create(ctx, taskType, cronSpec, payload)
}

// Show task
// GET /task/:id
func (c *Controller) Show(ctx context.Context, id uint) (*entity.Task, error) {
	return c.uc.Show(ctx, id)
}

// Edit returns a view for editing a task
// GET /task/:id/edit
func (c *Controller) Edit(ctx context.Context, id uint) (*entity.Task, error) {
	return c.uc.Edit(ctx, id)
}

// Update task
// PATCH /task/:id
func (c *Controller) Update(ctx context.Context, id uint, cronSpec string, payload string) error {
	return c.uc.Update(ctx, id, cronSpec, payload)
}

// Delete task
// DELETE /task/:id
func (c *Controller) Delete(ctx context.Context, id uint) error {
	return c.uc.Delete(ctx, id)
}
