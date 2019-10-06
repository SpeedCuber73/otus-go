package sheduler

import (
	"fmt"
	"testing"
	"time"
)

func makeTaskFunc(waitFor int, err error) func() error {
	return func() error {
		fmt.Println("start ", waitFor, " sec goroutine")
		time.Sleep(time.Duration(waitFor) * time.Second)
		return err
	}
}

const countTasks = 5

func TestSheduleTasks(t *testing.T) {
	tasks := make([]func() error, countTasks)
	for i := 0; i < countTasks; i++ {
		tasks = append(tasks, makeTaskFunc(i, nil))
	}

	SheduleTasks(tasks, 3, 0)
}
