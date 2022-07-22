package channel

import (
	"sync"
	"time"
)

func TestChannel() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 13
	}()
	n := <- ch1
	println(n)
}

func produce(ch chan<- int) {
	for i :=0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

func consume(ch <-chan int) {
	for n := range ch {
		println(n)
	}
}

func Main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		produce(ch)
		wg.Done()
	}()

	go func() {
		consume(ch)
		wg.Done()
	}()

	wg.Wait()
}


