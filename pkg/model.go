package pkg

import (
	"fmt"
	"net"
	"time"
)

type Task struct {
	Endpoint string
	Timeout  time.Duration
	Status   bool
	Error    error
}

func NewTaskWithOption(host, port string, timeout ...time.Duration) *Task {
	return NewTask(fmt.Sprintf("%s:%s", host, port), timeout...)
}

func NewTask(endpoint string, timeout ...time.Duration) *Task {
	task := &Task{
		Endpoint: endpoint,
		Timeout:  time.Second * 1,
	}
	if len(timeout) > 0 {
		task.Timeout = timeout[0]
	}
	return task
}

func (t *Task) Do() error {
	conn, err := net.DialTimeout("tcp", t.Endpoint, t.Timeout)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}
