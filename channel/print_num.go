package cha

import (
	"fmt"
	"sync"
)

//一个生产者，三个消费者，消费10个数字
var count int

func producer(ch chan int) {
	ch <- 1
}

func consumer(ch,nextCh chan int, wg *sync.WaitGroup) {
	for {
		<-ch
		if count >= 10 {
			nextCh <- 1
			wg.Done()
			return
		}
		count++
		fmt.Println(count)

		nextCh <- 1
	}
}

func TestCha() {
	var chs [] chan int
	for i := 0; i < 3; i++ {
		chs = append(chs, make(chan int,1))
	}
	wg := sync.WaitGroup{}
	wg.Add(3)
	go consumer(chs[0],chs[1], &wg)
	go consumer(chs[1],chs[2], &wg)
	go consumer(chs[2],chs[0], &wg)

	producer(chs[0])

	wg.Wait()
}
