package threadpoolgo

import (
	"sync"
)

type Pool struct {
	WorkerNum int

	TaskQueue chan Task

	dispatcher *dispatcher

	wg sync.WaitGroup
}
func NewPool(WorkerNum int, TaskQueueLen int) *Pool{
	TaskQueue :=make(chan Task, TaskQueueLen)
	workerQueue:=make(chan *worker, WorkerNum)
	return &Pool{
		TaskQueue:  TaskQueue,
		dispatcher: newDispatcher(workerQueue, TaskQueue),
	}
}

func(p *Pool)AddTask(task Task){
	p.TaskQueue <-task
}
func(p *Pool)Done(){
	p.wg.Done()
}
func(p *Pool)WaitTotal(total int){
	p.wg.Add(total)
}
func(p *Pool)Wait(){
	p.wg.Wait()
}
func(p *Pool)Release(){
	p.dispatcher.Stop <-struct{}{}
	<-p.dispatcher.Stop
}