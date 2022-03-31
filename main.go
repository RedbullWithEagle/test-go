package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/test-go/other"
)

func abc() (int,error){
	return 3,errors.New("abc")
}
func sum (a,b int) (err error){
	ret,err :=abc()
	fmt.Println(err)
	fmt.Println(ret)
	return
}
func main() {
	other.TestWithCar()
	sum(2,3)
	go func() {
		fmt.Println("hello goroutine")
	}()

	time.Sleep(1*time.Minute)
}
