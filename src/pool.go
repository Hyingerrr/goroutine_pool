package src

import "fmt"

// 协程池相关
type Pools struct {
	gCount     int        // goroutine count
	jobsQueue  chan *Task // 执行queue
	entryQueue chan *Task // 任务queue
}

func NewPool(cap int) *Pools {
	return &Pools{
		gCount:     cap,
		jobsQueue:  make(chan *Task),
		entryQueue: make(chan *Task),
	}
}

func (p *Pools) worker() {
	// worker 不断的从jobsQueue中拿任务 执行
	for task := range p.jobsQueue {
		err := task.Execute()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (p *Pools) Run() {
	// 开启gCount个goroutine，每个worker用一个goroutine承载
	for i := 0; i< p.gCount; i++ {
		go p.worker()
	}

	for t := range p.entryQueue {
		p.jobsQueue <- t
	}

	//close(p.entryQueue)
}