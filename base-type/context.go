package base_type

import (
	"context"
	"fmt"
	"time"
)

type otherContext struct {
	context.Context
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			fmt.Printf("%s is runnig,%d\n", name, int(time.Now().Unix()))
			time.Sleep(1 * time.Second)
		}
	}
}

func workWithValue(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			value := ctx.Value("key").(string)
			fmt.Printf("%s is runnig value=%s,%d\n", name, value, int(time.Now().Unix()))
			time.Sleep(1 * time.Second)
		}
	}
}

func testContext() {
	ctxa, cancel := context.WithCancel(context.Background())

	go work(ctxa, "work1")

	tm := time.Now().Add(3 * time.Second)
	ctxb, _ := context.WithDeadline(ctxa, tm)

	go work(ctxb, "work2")
	oc := otherContext{ctxb}
	ctxc := context.WithValue(oc, "key", "andes,pass from main")

	go workWithValue(ctxc, "work3")

	time.Sleep(10 * time.Second)
	cancel()

	time.Sleep(5 * time.Second)
	fmt.Println("main stop")
}
