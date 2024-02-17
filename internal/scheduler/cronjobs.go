package scheduler

import (
	"context"
	"sync"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/lohuza/relayer/internal/scheduler/tasks"
	"github.com/lohuza/relayer/pkg"
	"github.com/rs/zerolog/log"
)

type Scheduler struct {
	scheduler gocron.Scheduler
}

func NewScheduler(ctx context.Context, wg *sync.WaitGroup) *Scheduler {
	cronScheduler, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize cron")
	}

	scheduler := &Scheduler{
		scheduler: cronScheduler,
	}

	pkg.HandleShutdown(ctx, wg, func() {
		scheduler.scheduler.StopJobs()
	})

	return scheduler
}

func (s *Scheduler) Start() {
	_, err := s.scheduler.NewJob(gocron.DurationJob(30*time.Second), gocron.NewTask(tasks.FundAccountsTask))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start cron for funding accounts")
	}

	_, err = s.scheduler.NewJob(gocron.DurationJob(10*time.Second), gocron.NewTask(tasks.FetchGasPricesTask))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start cron for fetching gas prices")
	}

	s.scheduler.Start()
}
