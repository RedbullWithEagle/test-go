package base_type

import (
	"fmt"
	"sync"
)

func TestSyncMap() {
	var smp sync.Map

	// 数据写入
	smp.Store("name", "小红")
	smp.Store("age", 18)

	// 数据读取
	name, ok := smp.Load("name")
	if ok {
		fmt.Println(name)
	} else {
		fmt.Println("not found key-name")
	}

	age, ok := smp.Load("age")
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("not found key-age")
	}

	// 遍历
	smp.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	// 删除
	smp.Delete("age")
	age1, ok1 := smp.Load("age")
	fmt.Println("删除后的查询")
	fmt.Println(age1, ok1)

	// 读取或写入,存在就读取，不存在就写入
	v,ok2:=smp.LoadOrStore("age", 100)
	fmt.Println("LoadOrStore",v,ok2)
	age, ok = smp.Load("age")
	fmt.Println("不存在")
	fmt.Println(age, ok)

	smp.LoadOrStore("age", 99)
	age, ok = smp.Load("age")
	fmt.Println("存在")
	fmt.Println(age, ok)
}
