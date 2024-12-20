package scheduler

import (
	"log"

	"github.com/go-co-op/gocron/v2"
)

var (
	Scheduler gocron.Scheduler
)

func init() {
	var err error

	Scheduler, err = gocron.NewScheduler()
	if err != nil {
		log.Fatalf("Error during scheduler creation: %v", err)
	}

	Scheduler.Start()
}
