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
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//slice 是引用传递，会修改外面的值
	/*for k, v := range a {
		if k == 0 {
			a[k], a[k+1] = 100, 200
			fmt.Println(a)
		}

		a[k] = v + 100
	}*/
	sliceAddBatch(a, 10, 30)
	fmt.Println("test slice ")
	fmt.Println(a, len(a), cap(a))
}

func sliceAdd(arr []int) {
	arr[1] = 100

	fmt.Println("slice add ")
	fmt.Println(arr)
}

/********************************************
1.slice如何删除元素
2.slice当做函数参数
属于引用传递,其实也是值传递slice的结果，pointer，len，cap
如果函数内，修改slice长度，删除元素，外面是可以看到
新增元素，外面是看不到的
如果新增元素发生扩容，那么底层和之前的slice分离
*********************************************/
func sliceDelete(arr []int, index, count int) {
	arr = append(arr[:index], arr[index+count:]...)
	fmt.Println(arr, len(arr), cap(arr))
}

func sliceAddBatch(arr []int, index, count int) {
	for i := index; i < index+count; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr, len(arr), cap(arr))
}

/******************************************************
返回子切片导致的内存泄露
*****************************************************/
const TOTAL = 10000

func ReturnSubSlice1(begin, end int) []int {
	parent := make([]int, TOTAL)

	//只要child没有被GC回收，parent就一直得不到释放
	child := parent[begin:end]
	return child
}

/******************************************************
解决上面内存泄露的问题
新make一个数组，并且将值拷贝过去
*****************************************************/
func ReturnSubSlice2(begin, end int) []int {
	parent := make([]int, TOTAL)

	//只要child没有被GC回收，parent就一直得不到释放
	child := make([]int, end-begin)
	for i := begin; i < end; i++ {
		child[i-begin] = parent[i]
	}
	return child
}

func UseSubSlice() {
	s1 := ReturnSubSlice1(2, 7)
	fmt.Printf("init len :%d,cap:%d,TOTAL=%d\n", len(s1), cap(s1), TOTAL)

	for i := 0; i < 5; i++ {
		s1 = append(s1, i+5)
		fmt.Printf("len :%d,cap:%d\n", len(s1), cap(s1))
	}
	fmt.Println("================================")
	s2 := ReturnSubSlice2(2, 7)
	fmt.Printf("init len :%d,cap:%d,TOTAL=%d\n", len(s2), cap(s2), TOTAL)

	for i := 0; i < 5; i++ {
		s2 = append(s2, i+5)
		fmt.Printf("len :%d,cap:%d\n", len(s2), cap(s2))
	}
	fmt.Println("================================")

	//copy函数只能用于slice之间的复制
	//底层调用了memmove函数
	//src和dst，哪个短，就拷贝短的个数的数据
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := make([]int, len(s3))
	copy(s4, s3)
	s4[1] = 88
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Printf("s3 address:%p,%p\n",&s3[0],&s3)
	fmt.Printf("s4 address:%p,%p\n",&s4[0],&s4)

	/*var s5 []int
	copy(s5,s3)
	s5[0]=22
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Printf("s4 address:%p",&s4[0])
	fmt.Printf("s5 address:%p",&s5[0])*/

}
