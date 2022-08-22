package other

import (
	"fmt"
	"time"
)
func init(){
	fmt.Println("other.date init")
}
/*******************************************************
*birthDay2Time   将年月日转化为Time类型
 ******************************************************/
func birthDay2Time(bir int) time.Time {
	day := bir % 100
	month := (bir / 100) % 100
	year := (bir / 10000) % 10000
	fmt.Println(fmt.Sprintf("bir:%d      yeat:%d,mon:%d,day:%d", bir, year, month, day))
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func TestBirthDay() {
	bir1 := 19880334
	bir2 := 29991231
	bir3 := 19880101
	fmt.Println(birthDay2Time(bir1))
	fmt.Println(birthDay2Time(bir2))
	fmt.Println(birthDay2Time(bir3))
}
