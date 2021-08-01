package threadpoolgo


type worker struct{
	workerQueue chan *worker
	taskQueue chan Task
	stop chan struct{}
}
func newWorker(workerQueue chan *worker)*worker{
	return &worker{
		workerQueue: workerQueue,
		taskQueue: make(chan Task),
		stop: make(chan struct{}),
	}
}
func (w *worker) start()  {
	go func() {
		var task Task
		for{
			w.workerQueue<-w
			select{
				case task=<-w.taskQueue:
				task()
				case<-w.stop:
				w.stop<- struct{}{}
				return
			}
		}
	}()

}
