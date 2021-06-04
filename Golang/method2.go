package main

import (
   "fmt"
)

type Circle struct {
	radius float64
}


func (c Circle) area() float64{
	return 3.14 * c.radius * c.radius
}

func (c *Circle) area2() float64 {
	return 3.14 * (*c).radius * (*c).radius
}


type integer int

func (i integer)print() {
	fmt.Println("i = ",i)
}




type Student struct {
	Name string
	Age int
}

//给*Student实现String()方法
func (stu *Student) String() string {
	str := fmt.Sprintf("Name = %v Age = %v ",stu.Name,stu.Age)
	return str
}





func main() {
	var c Circle
	c.radius = 4.0
	res := c.area()
	fmt.Println("面积是 res = ：",res)

	res2 := (&c).area2()
	fmt.Println("res2 = ",res2)


	var i integer = 100
	i.print()


	//定义一个student变量
	stu := Student {
		Name : "tom",
		Age :20,
	}

	fmt.Println(&stu)

}
