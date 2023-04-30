package usecase

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
)

type SchedulerUC struct {
	repo SchedulerRepo
}

func NewSchedulerUC(repo SchedulerRepo) *SchedulerUC {
	return &SchedulerUC{repo: repo}
}

func (s *SchedulerUC) GetConfigs() ([]*asynq.PeriodicTaskConfig, error) {
	var tasks, err = s.repo.List(context.Background())
	if err != nil {
		log.Printf("Error getting: %+v\n", err)
		return nil, err
	}

	var (
		configs  = make([]*asynq.PeriodicTaskConfig, len(tasks))
		tConfigs = make([]asynq.PeriodicTaskConfig, len(tasks))
	)
	for i := 0; i < len(tasks); i++ {
		tConfigs[i].Cronspec = tasks[i].CronSpec
		tConfigs[i].Task = asynq.NewTask(tasks[i].Type, []byte(tasks[i].Payload))

		configs[i] = &tConfigs[i]
	}

	return configs, nil
}
