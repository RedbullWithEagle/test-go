package other

import (
	"fmt"
	"time"
)

/***************************************************************************
*继承
*一个结构体嵌到另一个结构体，称作组合
*匿名和组合的区别
*如果一个struct嵌套了另一个匿名结构体，那么这个结构可以直接访问匿名结构体的方法，从而实现继承
*如果一个struct嵌套了另一个【有名】的结构体，那么这个模式叫做组合
*如果一个struct嵌套了多个匿名结构体，那么这个结构可以直接访问多个匿名结构体的方法，从而实现多重继承
***************************************************************************/

type Car struct {
	weight int
	name   string
}

func (p *Car) Run() {
	fmt.Println("running")
}

type Bike struct {
	Car
	lunzi int
}
type Train struct {
	Car
}

func (p *Train) String() string {
	str := fmt.Sprintf("name=[%s] weight=[%d]", p.name, p.weight)
	return str
}

func TestWithCar() {
	var a Bike
	a.weight = 100
	a.name = "bike"
	a.lunzi = 2
	fmt.Println(a)
	a.Run()

	var b Train
	b.weight = 100
	b.name = "train"
	b.Run()
	fmt.Printf("%s", &b)
}


type Car1 struct {
	Name string
	Age  int
}

func (c *Car1) Set(name string, age int) {
	c.Name = name
	c.Age = age
}

type Car2 struct {
	Name string
}

//Go有匿名字段特性
type Train2 struct {
	Car1
	Car2
	createTime time.Time
	//count int   正常写法，Go的特性可以写成
	int
}

//给Train加方法，t指定接受变量的名字，变量可以叫this，t，p
func (t *Train2) Set(age int) {
	t.int = age
}

func main() {
	var train Train2
	train.int = 300 //这里用的匿名字段写法，给Age赋值
	//(&train).Set(1000)
	train.Car1.Set("huas", 100 )
	train.Car1.Name = "test" //这里Name必须得指定结构体
	fmt.Println(train)
}