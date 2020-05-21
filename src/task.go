package src

type Logger interface {
	Print(v ...interface{})
}

type Executor struct {
	log  Logger
	task *Task
	pool *Pools
}

func NewExecutor(task *Task, pool *Pools, logger Logger) *Executor {
	return &Executor{
		log:  logger,
		task: task,
		pool: pool,
	}
}

func (e Executor) Processor() {
	go func() {
		for  {
			if e.task == nil {
				close(e.pool.entryQueue)
			}

			e.pool.entryQueue <- e.task
		}
	}()

	e.pool.Run()
}

// *********** Task ************** //

type Task struct {
	f func() error
}

func NewTask(f func() error) *Task {
	return &Task{f:f}
}

func (t *Task) Execute() error {
	return t.f()
}