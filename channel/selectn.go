package channel

import (
	"fmt"
	"sync"
	"time"
)

// 一个package下不能有同名的函数，因为go是用package.name来调用函数的
func worker2(i int) {
	fmt.Printf("worker %d: is working...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d: works done...\n", i)
}

type signal2 struct {

}

func spawnGroup(f func(i int), num int, groupSignal <-chan signal2) <-chan signal2{
	c := make(chan signal2)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			<- groupSignal
			fmt.Printf("worker %d: start to work \n", i)
			f(i)
			wg.Done()
		}(i+1)
	}

	go func() {
		wg.Wait()
		c <- signal2{}
	}()
	return c
}

func SelectMain2() {
	//main goroutine 创建了一组 5 个 worker goroutine，这些 Goroutine 启动后会阻塞在名为 groupSignal 的无缓冲 channel 上。
	//main goroutine 通过close(groupSignal)向所有 worker goroutine 广播“开始工作”的信号，收到 groupSignal 后，
	//所有 worker goroutine 会“同时”开始工作，就像起跑线上的运动员听到了裁判员发出的起跑信号枪声
	fmt.Println("start a group of workers")
	groupSignal := make(chan signal2)
	c := spawnGroup(worker2, 5, groupSignal)
	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work")
	close(groupSignal)
	<-c
	fmt.Println("the group of workers work done")
}
