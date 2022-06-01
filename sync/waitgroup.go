package sync

import (
	"fmt"
	"sync"
	"time"
)

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