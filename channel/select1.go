package channel

import (
	"fmt"
	"time"
)

type signal struct {
}

func worker() {
	println("worker is working")
	time.Sleep(1 * time.Second)
}

// <-chan 代表只接收channel，表示从channel中接收
// 无缓冲 channel 用做信号传播，1:1 or 1:n
func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("worker start to work...")
		f()
		c <- signal{}
	}()
	return c
}

func SelectMain() {
	println("start a worker")
	// spawn 返回channel，用于新Goroutine的退出的通知信号，通知 SelectMain 的Goroutine上
	// SelectMain goroutine 在调用 spawn 函数后一直阻塞在对这个“通知信号”的接收动作上
	c := spawn(worker)
	<-c
	fmt.Println("worker work done")
}
