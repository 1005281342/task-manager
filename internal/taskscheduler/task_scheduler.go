package taskscheduler

import (
	"sync"

	"github.com/1005281342/task-manager/internal/config"
	"github.com/1005281342/task-manager/internal/controller/scheduler"
	"github.com/1005281342/task-manager/internal/controller/worker"
)

func Run(cfg config.Config) {

	var wg = &sync.WaitGroup{}
	var s = scheduler.Controller{}
	wg.Add(1)
	s.Load(cfg, wg)

	var w = worker.Controller{}
	wg.Add(1)
	w.Load(cfg, wg)

	wg.Wait()
}
