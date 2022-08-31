package base_type

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

/*****************************************************
time.After 使用局部变量会引起内存泄露
for select 每次检查，会生成新的对象
直到时间到了，被执行时，那个对象才会被gc
****************************************************/
func TestTimeAfter() {
	ch := make(chan string, 100)

	go func() {
		for {
			ch <- "asong"
		}
	}()

	go func() {
		// 开启pprof，监听请求
		ip := "127.0.0.1:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
		}
	}()

	idleDuration := 5 * time.Millisecond
	idleTimeout := time.NewTimer(idleDuration)
	defer idleTimeout.Stop()

	for {
		select {
		case <-ch:
			fmt.Printf("print:\n")
			// 这里如果换成time.After会有内存泄露
		case <-idleTimeout.C:
			fmt.Printf("after:\n")
		}
	}
}
