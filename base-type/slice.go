package base_type

import "fmt"

/*************************************************************
切片和数组区别
1.数组是固定长度，常量。切片长度是可以改变，所以是一个可变的数组
2.数组是值类型，数组传递是值。切片是引用类型，切片传递是指针
  切片作为函数参数，也会拷贝一份。如果只修改现有数据的值，外面的切片是可以看到这个修改
  如果长度增加，外面的切片是看不到新增的长度
3.数组不能使用append添加元素(因为数组是常量)，切片通过append添加元素
 ************************************************************/
func TestArray() {
	a := [3]int{1, 2, 3}

	//数组是值传递，这里的v修改后不会变
	for k, v := range a {
		/*if k == 0 {
			a[k], a[k+1] = 100, 200
			fmt.Println(a)
		}*/
		a[1] = 300
		a[k] = v + 100
	}
	fmt.Println(a)
}

func TesSlice() {
	a := []int{1, 2, 3}

	//slice 是引用传递，会修改外面的值
	/*for k, v := range a {
		if k == 0 {
			a[k], a[k+1] = 100, 200
			fmt.Println(a)
		}

		a[k] = v + 100
	}*/
	sliceAdd(a)
	fmt.Println("test slice ")
	fmt.Println(a)
}

func sliceAdd(arr []int){
	arr[1] = 100

	fmt.Println("slice add ")
	fmt.Println(arr)
}