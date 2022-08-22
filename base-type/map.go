package base_type

import "fmt"

/*************************************************
1. value,ok := map[key]  如果元素不存在，ok=false
 ************************************************/
func TestOKMap() {
	mMap := make(map[string]int, 0)
	mMap["liu"] = 123

	//map如果元素不存在，ok=false；反向判断的时候需要注意
	if value, ok := mMap["liu"]; ok {
		fmt.Println(value)
	}

	value, ok := mMap["uuu"]
	fmt.Println(value, ok) //0,false
}

type student struct {
	Name string
	Age  int
}

func TestPointerMap() {
	m := make(map[string]student)

	stus := []student{
		{Name: "liu", Age: 34},
		{Name: "zhao", Age: 25},
		{Name: "wang", Age: 16},
	}

	//如果切片是指针类型，则打印的都是一个地址
	//即最后赋值的地址
	for _, v := range stus {
		m[v.Name] = v
	}
	fmt.Println(m)
}
