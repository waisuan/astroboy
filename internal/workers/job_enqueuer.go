package workers

import (
	"astroboy/internal/dependencies"
	"github.com/gocraft/work"
)

type Context struct{}

type JobEnqueuer struct {
	Pool *work.WorkerPool
}

func NewJobEnqueuer(deps *dependencies.Dependencies) *JobEnqueuer {
	pool := work.NewWorkerPool(Context{}, 1, "dummy_namespace", deps.CacheCli.Pool())

	registerSchedule(pool, deps.Config)
	registerJobs(pool)

	return &JobEnqueuer{
		Pool: pool,
	}
}

func registerSchedule(pool *work.WorkerPool, cfg *dependencies.Config) {
	pool.PeriodicallyEnqueue(cfg.FakeDataPublisherCron, "publish_fake_data")
}

func registerJobs(pool *work.WorkerPool) {
	pool.Job("publish_fake_data", (*Context).PublishFakeData)
}
