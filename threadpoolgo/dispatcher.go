package threadpoolgo


type dispatcher struct {
	WorkerQueue chan *worker
	TaskQueue   chan Task
	Stop        chan struct{}
}
func newDispatcher(WorkerQueue chan *worker, TaskQueue chan Task)*dispatcher{
	dispatcher:=&dispatcher{
		WorkerQueue: WorkerQueue,
		TaskQueue:   TaskQueue,
		Stop:        make(chan struct{}),
	}
	for i := 0; i < cap(dispatcher.WorkerQueue); i++ {
		worker:=newWorker(dispatcher.WorkerQueue)
		worker.start()

	}
	go dispatcher.dispatcher()
	return dispatcher
}
func(d *dispatcher)dispatcher(){
	for{
		select {
		case task:=<-d.TaskQueue:
		worker:=<-d.WorkerQueue
		worker.TaskQueue <-task
		case <-d.Stop:
			for i := 0; i < cap(d.WorkerQueue); i++ {
				worker:=<-d.WorkerQueue
				worker.Stop <-struct{}{}
					<-worker.Stop
			}
			d.Stop <- struct{}{}
			return
		}
	}
}
