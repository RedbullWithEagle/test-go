package main

import (
	"time"

	"github.com/test-go/sync"
)

func main() {
	sync.TestWaitGroupWait()
	time.Sleep(5 * time.Second)
}
