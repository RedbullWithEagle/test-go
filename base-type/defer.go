package base_type

import "fmt"

func TestDefer(){
	 val := "111"
	 defer func() {
	 	fmt.Println(val)
	 }()

	 defer fmt.Println(val)

	 val = "222"
	 defer func(val string) {
	 	fmt.Println(val)
	 }(val)


	val = "333"
	defer func() {
		fmt.Println(val)
	}()

}

func ExecutePanic(){
	defer fmt.Println("defer func")
	defer func() {
		fmt.Println("recover func")
		if errMsg := recover();errMsg !=nil{
			fmt.Println(errMsg)
		}
		fmt.Println("This is   recover func situation")
	}()
	panic("This is Panic situation")
	fmt.Println("The function executes Completely")
}

var g = 100

/**********************************************************
defer return 执行顺序
1.将返回值保存在栈上
2.执行defer函数
3.函数返回
f()执行后， r=200，g=100
f1()执行后，r=100,g=200
************************************************************/
func f()(r int){
	r = g
	defer func() {
		r = 200
	}()

	r = 0
	return r
}

func f1() int{
	defer func() {
		g = 200
	}()

	return g
}

func Testf(){
	i := f1()

	fmt.Printf("i:=%d,g:= %d\n",i,g)
}