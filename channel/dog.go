package cha

import (
	"fmt"
	"sync"
	"time"
)

const Counter = 10

/*********************************************************
1.采用无缓冲的channel来顺序打印dog，fish，cat 10次
2.需要注意，最后打印cat的时候，不要给在dogChan发送数据，因为此时dogChan已经关闭
********************************************************/
func Dog(dogChan, fishChan chan int, wg *sync.WaitGroup) {
	defer close(dogChan)
	count := 0
	for {
		if count >= Counter {
			wg.Done()
			return
		}

		if v := <-dogChan; v == 1 {
			count++
			fmt.Printf("dog:%d\n", count)
			if count <= Counter {
				fishChan <- 1
			}
		}
	}
}

func Fish(fishChan, catChan chan int, wg *sync.WaitGroup) {
	defer close(fishChan)
	count := 0
	for {
		if count >= Counter {
			wg.Done()
			return
		}

		if v := <-fishChan; v == 1 {
			count++
			fmt.Printf("fish:%d\n", count)
			if count <= Counter {
				catChan <- 1
			}

		}
	}
}

func Cat(catChan, dogChan chan int, wg *sync.WaitGroup) {
	defer close(catChan)
	count := 0
	for {
		if count >= Counter {
			wg.Done()
			return
		}

		if v := <-catChan; v == 1 {
			count++
			fmt.Printf("cat:%d\n", count)
			if count < Counter {
				dogChan <- 1
			}
		}
	}
}

func TestDog() {
	wg := sync.WaitGroup{}
	dogCh := make(chan int)
	fishCh := make(chan int)
	catCh := make(chan int)
	wg.Add(3)

	go Dog(dogCh, fishCh, &wg)
	go Fish(fishCh, catCh, &wg)
	go Cat(catCh, dogCh, &wg)

	dogCh <- 1
	wg.Wait()
}

/************************************************************
如何限制并发协程的数量？
控制channel缓冲的数量来实现
 **********************************************************/
func job(i int, limitChan chan struct{}, group *sync.WaitGroup) {
	limitChan <- struct{}{}
	defer func() {
		<-limitChan
	}()

	time.Sleep(1 * time.Second)
	fmt.Printf("任务id:%d已完成，当前协程数:%d\n", i, len(limitChan))
	group.Done()

}

// limitChan限制协程数
func RunGoroutine() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	limitChan := make(chan struct{}, 10) // 最大协程数限制为10个
	for i := 0; i < 100; i++ {
		go job(i, limitChan, &wg)
	}
	wg.Wait()
}
