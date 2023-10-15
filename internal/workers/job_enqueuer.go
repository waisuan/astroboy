package workers

import (
	"astroboy/internal/dependencies"
	"github.com/gocraft/work"
	"log"
)

type Context struct{}

type JobEnqueuer struct {
	Pool *work.WorkerPool
}

func NewJobEnqueuer(deps *dependencies.Dependencies) *JobEnqueuer {
	pool := work.NewWorkerPool(Context{}, 1, "dummy_namespace", deps.CacheCli.Pool())

	registerJobs(pool)

	return &JobEnqueuer{
		Pool: pool,
	}
}

func registerJobs(pool *work.WorkerPool) {
	pool.PeriodicallyEnqueue("0 */3 * * * *", "publish_fake_data")
	pool.Job("publish_fake_data", (*Context).PublishFakeData)
}

func (c *Context) PublishFakeData(job *work.Job) error {
	log.Println("Running publish_fake_data job...")
	return nil
}
