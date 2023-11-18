package workerpool

import (
	"context"
	"sync"
)

// Task is a function that performs a job and returns an error if it fails.
type Task func(ctx context.Context) error

// WorkerPool interface defines the worker pool's functionality.
type WorkerPool interface {
	Submit(task Task)
	Run(ctx context.Context) error
}

// workerPool is the implementation of WorkerPool.
type workerPool struct {
	tasks chan Task
	wg    sync.WaitGroup
}

// New creates a new WorkerPool with the given number of workers.
func New(workerCount int) WorkerPool {
	return &workerPool{
		tasks: make(chan Task, workerCount),
	}
}

// Submit adds a new task to the pool.
func (p *workerPool) Submit(task Task) {
	p.wg.Add(1)
	p.tasks <- task
}

// Run starts the worker pool and waits for all tasks to complete.
func (p *workerPool) Run(ctx context.Context) error {
	var err error
	for i := 0; i < cap(p.tasks); i++ {
		go func() {
			for task := range p.tasks {
				if e := task(ctx); e != nil {
					err = e
				}
				p.wg.Done()
			}
		}()
	}

	p.wg.Wait()
	close(p.tasks)
	return err
}
