package threadpoolgo


type worker struct{
	WorkerQueue chan *worker
	TaskQueue   chan Task
	Stop        chan struct{}
}
func newWorker(WorkerQueue chan *worker)*worker{
	return &worker{
		WorkerQueue: WorkerQueue,
		TaskQueue:   make(chan Task),
		Stop:        make(chan struct{}),
	}
}
func (w *worker) start()  {
	go func() {
		var task Task
		for{
			w.WorkerQueue <-w
			select{
				case task=<-w.TaskQueue:
				task()
				case<-w.Stop:
				w.Stop <- struct{}{}
				return
			}
		}
	}()

}
