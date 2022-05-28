package pkg

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

const TEMPLATE = "%d.%d.%d.%d:%d"

// 端口扫描
func StartPort(host string) {
	taskChan := make(chan Task, 10)
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
		}
		// else {
		// 	color.Red("[失败] 扫描地址：%s; 失败原因：%s\n", result.Endpoint, result.Error.Error())
		// }
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
