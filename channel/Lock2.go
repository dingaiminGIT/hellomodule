package channel

import (
	"fmt"
	"sync"
)

type counter2 struct {
	c chan int
	i int
}

// 函数
func NewCounter() *counter2 {
	cter := &counter2{
		// 无缓冲的channel，会发生同步阻塞
		c: make(chan int),
	}
	// 在这个goroutine中进行计数，但因为是无缓冲的channel，只有另一个goroutine从channel中把数据接收走，才能再次产生数据写到channel中
	go func() {
		for {
			cter.i++
			cter.c <- cter.i
		}
	}()
	return cter
}

// 方法
func (cter *counter2) Increase() int {
	// 调用Increase想从无缓冲的channel中读取数据即从channel接收
	return <- cter.c
}

func Lock2Main() {
	cter := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := cter.Increase()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}


