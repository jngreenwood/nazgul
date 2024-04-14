package worker

import (
	"context"
	"fmt"
)

//task is a type of work we can do, we then inherit from this type to make our jobs

type Task interface {
	Run(ctx context.Context)
	Cancel()
	Name() string
}

type JobScheduler struct {
	Queue      chan Task
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewJobScheduler(ctx context.Context) *JobScheduler {
	ctx, cancelFunc := context.WithCancel(ctx)
	return &JobScheduler{
		Queue:      make(chan Task),
		ctx:        ctx,
		cancelFunc: cancelFunc,
	}
}

func (js *JobScheduler) Start() {
	go func() {
		for {
			select {
			case task := <-js.Queue:
				task.Run(js.ctx)
			case <-js.ctx.Done():
				// Handle the cancellation
				fmt.Println("Scheduler stopped")
				return
			}
		}
	}()
}

func (js *JobScheduler) Stop() {
	fmt.Println("stopping jobs")
	// js.cancelFunc() // Trigger cancellation
	// Wait and clean up tasks
	for task := range js.Queue {
		task.Cancel()
	}
}

func (js *JobScheduler) AddTask(t Task) error {
	select {
	case js.Queue <- t:
	case <-js.ctx.Done():
		return fmt.Errorf("scheduler is stopped")
	}
	return nil
}
