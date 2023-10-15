package workers

import (
	"astroboy/internal/dependencies"
	"github.com/gocraft/work"
)

type Context struct{}

type JobEnqueuer struct {
	Pool *work.WorkerPool
}

func NewJobEnqueuer(cache dependencies.ICache) *JobEnqueuer {
	pool := work.NewWorkerPool(Context{}, 1, "dummy_namespace", cache.Pool())

	registerSchedule(pool)
	registerJobs(pool)

	return &JobEnqueuer{
		Pool: pool,
	}
}

func registerSchedule(pool *work.WorkerPool) {
	pool.PeriodicallyEnqueue("0 */3 * * * *", "publish_fake_data")
}

func registerJobs(pool *work.WorkerPool) {
	pool.Job("publish_fake_data", (*Context).PublishFakeData)
}
