package sync

import (
	"fmt"
	"sync"
	"time"
)

/******************************************************
一个 WaitGroup 对象可以等待一组协程结束
******************************************************/
func TestWaitGroup(){
	tasks  := []func(){
		func() { time.Sleep(time.Second); fmt.Println("1 sec later") },
		func() { time.Sleep(time.Second *  2); fmt.Println("2 sec later") },
	}

	var wg sync.WaitGroup // 1-1
	wg.Add(len(tasks))    // 1-2
	for _, task := range tasks {
		task  := task
		go func() {       // 1-3-1
			defer wg.Done() // 1-3-2
			task()          // 1-3-3
		}()               // 1-3-1
	}
	wg.Wait()           // 1-4
	fmt.Println("exit")
}

/************************************************************
1.wg add 之后
2.可以多次 wg.Wait，只要没有done
3.wg.done后，可以唤起所有的wg.Wait
************************************************************/
func TestWaitGroupWait() {
	var wg sync.WaitGroup
	fmt.Printf("test wait group:%d\n",time.Now().Unix())

	wg.Add(1)
	for i := 0; i < 10; i++ {
		go func() {
			wg.Wait()
			fmt.Printf("wg wait:%d\n",time.Now().Unix())
		}()
	}
	time.Sleep(5 * time.Second)
	wg.Done()
}