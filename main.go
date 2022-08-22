package main

import (
	"fmt"
	"time"

	base_type "github.com/test-go/base-type"
)

func main() {
	if err :=base_type.TestPanicError();err !=nil{
		fmt.Println(err)
	}

	time.Sleep(5*time.Second)
}
