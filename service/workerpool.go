package service

func NewWorkerPool(maxWorkerCount int) *workerPool {
	return &workerPool{
		JobCount: maxWorkerCount,
		Task:     make(chan func()),
	}
}

type workerPool struct {
	JobCount int
	Task     chan func()
}

func (wp *workerPool) Run() {
	for {
		for i := 0; i < wp.JobCount; i++ {
			go func(workerID int) {
				for {
					select {
					case task := <-wp.Task:
						task()
					}
				}
			}(i + 1)
		}
	}
}

func (wp *workerPool) Add(task func()) {
	wp.Task <- task
}
