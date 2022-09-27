package base_type

import (
	"errors"
	"fmt"
)

func getInteger() error {
	return errors.New("not integer err ")
}

func getString() error {
	return nil
}

func getABC() error {
	return errors.New("not abc err ")
}

/*************************************
*TestError 变量作用域
* :=声明一个新的变量，作用域在{}内
*结果如下：
*not integer err
*error test
*not abc err
*error test
************************************/
func TestError() error {
	var err error
	if err := getInteger(); err != nil {
		fmt.Println(err)
	}

	err = errors.New("error test")
	if err := getString(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(err)
	if err := getABC(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)
	return err
}

func TestFuncError() error{
	var err error
	defer func(err error) {
		err = errors.New("TestFuncError")
	}(err)

	return err
}

/**********************************************************
panic和recover使用注意事项
1.recover必须在defer函数中使用，但是不能被 defer 直接调用
   defer recover不能恢复
2.recover只能恢复同一个协程中的panic，
所以必须与可能发生panic的协程在同一个协程中才生效
3.如果返回error不声明，在内部声明，defer里面的赋值是不会返回的
************************************************************/
func TestPanicError() (err error){
	defer func() {
		if r:= recover(); r!=nil{
			err = errors.New(fmt.Sprintf("recover:%s",r))
		}
	}()

	//如果这里协程方式启动raisePanic，则不会捕获到
	raisePanic()
	return err
}

func raisePanic(){
	panic("have a painc")
}

type Error struct {
	Code   int
	Reason string
}

func (e Error) Error() string {
	return fmt.Sprintf("%v:%v", e.Code, e.Reason)
}

func ops() error {
	var err *Error = &Error{
		Code: 0,
		Reason: "",
	}
	return err
}

func TestOPSError()  {
	err := ops()
	fmt.Println(err)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("success")
	}

}
