package sheduler

import (
	"sync"
)

func worker(wg *sync.WaitGroup, input <-chan func() error, errCh chan<- error) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-input:
			if !ok {
				return
			}
			err := task()
			if err != nil {
				errCh <- err
			}
		}
	}
}

func errorNotifier(errCh <-chan error, stopCh chan<- struct{}, maxErrors int) {
	for range errCh {
		maxErrors--
		if maxErrors < 0 {
			stopCh <- struct{}{}
			return
		}
	}
}

func producer(wg *sync.WaitGroup, tasks []func() error, input chan<- func() error, errCh <-chan error, maxErrors int) {
	defer close(input)

	stopCh := make(chan struct{})
	go errorNotifier(errCh, stopCh, maxErrors)

	for i := 0; i < len(tasks); i++ {
		select {
		case input <- tasks[i]:
		case <-stopCh:
			return
		}
	}
}

// SheduleTasks shedule tasks through goroutines
func SheduleTasks(tasks []func() error, maxSimultaneusly, maxErrors int) {
	tasksCh := make(chan func() error)
	errCh := make(chan error)

	var wg sync.WaitGroup
	wg.Add(maxSimultaneusly + 1)

	for i := 0; i < maxSimultaneusly; i++ {
		go worker(&wg, tasksCh, errCh)
	}

	go producer(&wg, tasks, tasksCh, errCh, maxErrors)
	wg.Wait()
}
