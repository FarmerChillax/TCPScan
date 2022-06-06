package pkg

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

const (
	HOST_TEMPLATE     = "%v.%v.%v.%v"
	ENDPOINT_TEMPLATE = "%v:%v"
)

func Start(segments []string, startPort, endPort int64, workers int) {
	taskChan := make(chan Task, workers)
	resultChan := make(chan Task, 10)
	done := make(chan struct{})

	// 开启 worker
	for i := 0; i < workers; i++ {
		go worker(taskChan, resultChan, done)
	}
	// 获取 worker 完成状态
	go CloseResult(done, resultChan, workers)
	// 推送任务
	go PushTasks(segments, taskChan)

	ProcessResult(resultChan)
}

func PushTasks(segments []string, taskChan chan Task, timeout ...time.Duration) {
	switch len(segments) {
	case 3:
		for i := 1; i < 255; i++ {
			host := fmt.Sprintf(HOST_TEMPLATE, segments[0], segments[1], segments[2], i)
			startPort(host, 1, 65535, taskChan, timeout...)
		}
	}
}

func startPort(host string, start, end int, taskChan chan Task, timeout ...time.Duration) {
	for port := start; port < end; port++ {
		endpoint := fmt.Sprintf(ENDPOINT_TEMPLATE, host, port)
		InitTask(endpoint, taskChan, timeout...)
	}
}

// 端口扫描
func StartPort(host string) {
	taskChan := make(chan Task, 200)
	result := make(chan Task, 10)
	done := make(chan struct{})
	workers := cap(taskChan)
	// 开启 worker 处理任务
	for i := 0; i < workers; i++ {
		go worker(taskChan, result, done)
	}

	// 获取 worker 完成状态
	go CloseResult(done, result, workers)

	// 开启结果收集进程
	go ProcessResult(result)

	go func(start, end int) {
		for i := start; i < end; i++ {
			endpoint := fmt.Sprintf(host, i)
			InitTask(endpoint, taskChan)
		}
		close(taskChan)
	}(1, 65535)
	ProcessResult(result)
}

func CloseResult(done chan struct{}, resule chan Task, workers int) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(done)
	close(resule)
}

func ProcessResult(results chan Task) {
	for result := range results {
		if result.Status {
			color.Green("[成功] 扫描地址：%s\n", result.Endpoint)
		} else {
			color.Red("[失败] 扫描地址：%s; 失败原因：%s\n", result.Endpoint, result.Error.Error())
		}
	}
}

func InitTask(endpoint string, taskChan chan<- Task, timeout ...time.Duration) {
	task := NewTask(endpoint, timeout...)
	taskChan <- *task
}

func worker(tasks, result chan Task, done chan struct{}) {
	for task := range tasks {
		if task.Error = task.Do(); task.Error == nil {
			task.Status = true
		}
		result <- task
	}
	done <- struct{}{}
}
