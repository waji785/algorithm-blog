package threadpoolgo

import (
	"sync"
)

type Pool struct {
	workerNum int

	taskQueue chan Task

	dispatcher *dispatcher

	wg sync.WaitGroup
}
func NewPool(workerNum int,taskQueueLen int) *Pool{
	taskQueue:=make(chan Task,taskQueueLen)
	workerQueue:=make(chan *worker,workerNum)
	return &Pool{
		taskQueue: taskQueue,
		dispatcher:newDispatcher(workerQueue,taskQueue),
	}
}

func(p *Pool)AddTask(task Task){
	p.taskQueue<-task
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
	p.dispatcher.stop<-struct{}{}
	<-p.dispatcher.stop
}