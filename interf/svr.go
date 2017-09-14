package interf

type AsyncTaskSvr struct {
	tasks []ITask
}

func NewAsyncTaskSvr() *AsyncTaskSvr {
	return new(AsyncTaskSvr)
}

func (t *AsyncTaskSvr) AddTask(task ITask) {
	t.tasks = append(t.tasks, task)
}

func (t *AsyncTaskSvr) Run() {
	for _, t := range t.tasks {
		go t.Run()
	}

	select {}
}
