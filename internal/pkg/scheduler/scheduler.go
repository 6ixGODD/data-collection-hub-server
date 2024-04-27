package scheduler

// TODO:
// 	1. Tasks for clear logger cache in redis and write to mongodb
// 	2. Tasks for clear query cache
// 	3. Tasks for clear zap logger cache
// 	4. Based on cron expression

type Tasks interface {
	ClearRedisLogCache()
	ClearQueryCache()
	SyncLogger()
}

type TasksImpl struct {
}

func NewScheduler() *TasksImpl {
	return &TasksImpl{}
}

func (t *TasksImpl) ClearRedisLogCache() {
	// TODO
}

func (t *TasksImpl) ClearQueryCache() {
	// TODO
}

func (t *TasksImpl) SyncLogger() {
	// TODO
}
