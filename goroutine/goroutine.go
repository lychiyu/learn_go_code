package main

import (
	"fmt"
	"time"
)

impomt"
	"time"
)

/**
协程

轻量级线程
非抢占式多任务处理，由协程主动交出控制权
编译器\解释器\虚拟机层级的多任务而非系统层级
多个协程可能在一个或多个线程上运行

goroutine
调度器会在合适的点进行协程的切换(不保证切换)，如：
i/o、select、channel、等待锁、函数调用(有时)、runtime.Gosched()
 */
func main() {
	for i := 0; i < 10000; i++ {
		go func(i int) {
			fmt.Printf("hello %d \n", i)
		}(i)
	}
	time.Sleep(time.Millisecond)
}
