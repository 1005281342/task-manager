package scheduler

import (
	"log"
	"sync"
	"time"

	"github.com/hibiken/asynq"

	"github.com/1005281342/task-manager/internal/config"
	"github.com/1005281342/task-manager/internal/entity"
	"github.com/1005281342/task-manager/internal/usecase"
	"github.com/1005281342/task-manager/internal/usecase/repo"
	"github.com/1005281342/task-manager/pkg/db"
)

var once sync.Once

// Controller for scheduler
type Controller struct{}

func New(cfg config.Config, db db.Connection) Controller {
	once.Do(func() {
		var err error
		if err = db.GetDB().AutoMigrate(&entity.Task{}); err != nil {
			log.Fatal(err)
		}

		var mgr *asynq.PeriodicTaskManager
		if mgr, err = asynq.NewPeriodicTaskManager(
			asynq.PeriodicTaskManagerOpts{
				RedisConnOpt:               asynq.RedisClientOpt{Addr: cfg.Redis.Addr},
				PeriodicTaskConfigProvider: usecase.NewSchedulerUC(repo.NewTaskRepo(db.GetDB())),
				// this provider object is the interface to your config source
				SyncInterval: 10 * time.Second, // this field specifies how often sync should happen
			}); err != nil {
			log.Fatal(err)
		}

		go func() {
			// start the scheduler
			log.Println("starting scheduler")
			if err = mgr.Run(); err != nil {
				log.Fatal(err)
			}
			log.Println("close scheduler")
		}()

		return
	})

	return Controller{}
}
