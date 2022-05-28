package main

import (
	"TCPScan/pkg"
	"flag"
	"fmt"
	"time"

	"github.com/pkg/profile"
)

var TIMEOUT = 1 * time.Second

func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	// 加载环境变量
	// ...
	host := flag.String("host", "127.0.0.1", "目标主机")
	flag.Parse()
	// args := flag.Args()
	// fmt.Println(args)
	// fmt.Println(*host)

	// 得到A-D网段
	// segmentList := []string{"10", "9", "10"}
	// 开始创建扫描任务
	// ...
	address := fmt.Sprintf("%s:%%d", *host)
	pkg.StartPort(address)
}

// func create(segments []string) {
// 	switch len(segments) {
// 	case 3:

// 	}
// }
