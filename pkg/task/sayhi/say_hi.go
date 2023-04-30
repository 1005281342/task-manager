package sayhi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"

	"github.com/1005281342/task-manager/internal/entity"
)

// Processor implements asynq.Handler interface.
type Processor struct {
	// ... fields for struct
}

func (processor *Processor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	log.Println(string(t.Payload()))
	var p entity.SayHi
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		log.Printf("Unmarshal failed: %+v\n", err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("hi %s \n", p.Name)
	return nil
}

func NewProcessor() *Processor {
	return &Processor{}
}
