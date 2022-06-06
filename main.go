package main

import (
	"TCPScan/pkg"
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/pkg/profile"
)

var TIMEOUT = 1 * time.Second

func main() {
	go func() {
		http.ListenAndServe(":1234", nil)
	}()
	defer profile.Start().Stop()
	// defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	// 加载环境变量
	// ...
	host := flag.String("host", "127.0.0.1", "目标主机")
	// segA := flag.String("A", "127", "A网段")
	// segB := flag.String("B", "0", "B")
	// segC := flag.String("C", "0", "C网段")
	// workers := flag.Int("worker", 20, "goroutine number.")
	flag.Parse()
	// segments := []string{*segA, *segB, *segC}
	// fmt.Println(segments, *workers)
	// args := flag.Args()

	// 得到A-D网段
	// segmentList := []string{"10", "9", "10"}
	// 开始创建扫描任务
	// ...
	address := fmt.Sprintf("%s:%%d", *host)
	pkg.StartPort(address)
	// pkg.Start(segments, *workers)
}
