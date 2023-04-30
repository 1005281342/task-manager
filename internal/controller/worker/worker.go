package worker

import (
	"log"
	"sync"

	"github.com/hibiken/asynq"

	"github.com/1005281342/task-manager/internal/config"
	"github.com/1005281342/task-manager/pkg/task"
	"github.com/1005281342/task-manager/pkg/task/sayhi"
)

// Controller for workers
type Controller struct {
	once sync.Once
}

func (c *Controller) Load(cfg config.Config, wg *sync.WaitGroup) {
	c.once.Do(func() {
		var srv = asynq.NewServer(
			asynq.RedisClientOpt{Addr: cfg.Redis.Addr},
			asynq.Config{
				// Specify how many concurrent workers to use
				Concurrency: 10,
				// Optionally specify multiple queues with different priority.
				Queues: map[string]int{
					"critical": 6,
					"default":  3,
					"low":      1,
				},
				// See the godoc for other configuration options
			},
		)

		// mux maps a type to a handler
		var mux = asynq.NewServeMux()
		mux.Handle(task.TypeSayHi, sayhi.NewProcessor())
		// ...register other handlers...

		go func() {
			defer wg.Done()
			log.Println("starting worker")
			if err := srv.Run(mux); err != nil {
				log.Fatalf("could not run server: %v", err)
			}
			log.Println("close worker")
		}()
	})
}
