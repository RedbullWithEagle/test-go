package channel

import (
	"fmt"
	"time"
)

/*********************************************************
*1.time.Tick与time.After有着本质的不通
*time.After并不会定时发送数据到通道中，而只是在时间到了后发送一次数据
*time.Tick不同，由于tick在for循环内部，因此其不重置，只会累积
*如果在time.After的时间到达之前，其他通道已经执行好了，那么time.After将不会得到执行
**********************************************************/
func Timer() {
	c := make(chan int, 1)
	tick := time.Tick(time.Second)

	for {
		select {
		case <-c:
			fmt.Println("random 01")
		case <-tick:
			fmt.Println("tick")
		case <-time.After(500 * time.Millisecond):
			fmt.Println("timeout")
		}
	}
}
