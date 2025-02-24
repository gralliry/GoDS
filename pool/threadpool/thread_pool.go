package threadpool

import (
	"sync"
)

type ThreadPool struct {
	wg    sync.WaitGroup
	once  sync.Once
	tasks chan func()
	stop  chan struct{}
}

func New(size int) *ThreadPool {
	tp := &ThreadPool{
		tasks: make(chan func(), size),
		stop:  make(chan struct{}),
	}
	return tp
}

func (p *ThreadPool) Wait() {
	p.wg.Wait()
}

func (p *ThreadPool) Stop() {
	p.once.Do(func() {
		close(p.tasks)
	})
}

func (p *ThreadPool) Close() {
	p.once.Do(func() {
		close(p.stop)
	})
}

func (p *ThreadPool) Start(worker int) {
	p.wg.Add(worker)
	for i := 0; i < worker; i++ {
		go p.execute()
	}
}

func (p *ThreadPool) Submit(task func()) bool {
	select {
	case p.tasks <- task:
		return true
	default:
		return false
	}

}

func (p *ThreadPool) execute() {
	defer p.wg.Done()
	for {
		select {
		case task, ok := <-p.tasks:
			if !ok {
				return
			}
			task()
		case <-p.stop:
			return
		}
	}
}
