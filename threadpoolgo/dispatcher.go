package threadpoolgo


type dispatcher struct {
	workerQueue chan *worker
	taskQueue chan Task
	stop chan struct{}
}
func newDispatcher(workerQueue chan *worker,taskQueue chan Task)*dispatcher{
	dispatcher:=&dispatcher{
		workerQueue: workerQueue,
		taskQueue: taskQueue,
		stop:make(chan struct{}),
	}
	for i := 0; i < cap(dispatcher.workerQueue); i++ {
		worker:=newWorker(dispatcher.workerQueue)
		worker.start()

	}
	go dispatcher.dispatcher()
	return dispatcher
}
func(d *dispatcher)dispatcher(){
	for{
		select {
		case task:=<-d.taskQueue:
		worker:=<-d.workerQueue
		worker.taskQueue<-task
		case <-d.stop:
			for i := 0; i < cap(d.workerQueue); i++ {
				worker:=<-d.workerQueue
				worker.stop<-struct{}{}
					<-worker.stop
			}
			d.stop<- struct{}{}
			return
		}
	}
}
