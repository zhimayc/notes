package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v \n",stu.Name,stu.Age,stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}


//小学生
type Pupil struct {
	Student  //嵌入了Student匿名结构体

}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试。。。")
}

//大学生
type Graduate struct {
	Student  //嵌入了Student匿名结构体
}

func (g *Graduate) testing() {
	fmt.Println("大学生正在考试。。。")
}



type A struct {
	Name string
	Age string
}

type B struct {
	Name string
	Age string
}

type C struct {
	a A		//有名结构体
	B  //匿名结构体
}

type Goods struct {
	Name string
	Price float64
}

type Brand struct {
	Name string
	Address string
}

type TV struct {
	Goods
	Brand
}



type GrandParent struct {
	Name string
	Age int
}

type Parent struct {
	GrandParent
	Gender string
}

type Child struct {
	Parent
	Salary float64
}





func main() {

	//当我们对结构体嵌入了匿名结构体使用方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "tom~"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(90)
	pupil.Student.ShowInfo()


	graduate := &Graduate{}
	graduate.Student.Name = "alice~"
	graduate.Student.Age = 18
	graduate.testing()
	graduate.Student.SetScore(95)
	graduate.Student.ShowInfo()


	//当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问，如果希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分
	//结构体嵌入两个（或多个）匿名结构体，如两个匿名结构体有相同的字段和方法（同时结构体本身没有同名的字段和方法），在访问时，就必须明确指定匿名结构体名字，否则编译报错
	//如果结构体中包含有名结构体，则访问有名结构体的字段时，就必须带上有名结构体的名字 例如 c.a.Name
	fmt.Println("-------------------------")
	graduate.Name = "Lucy"
	graduate.Age = 20
	graduate.SetScore(100)
	graduate.ShowInfo()

	//嵌套匿名结构体后，也可以在创建结构体变量（实例）时，直接指定各个匿名结构体字段的值
	tv := TV{ Goods{"电视机001",5999},Brand{"海尔","山东"} }
	fmt.Println("tv",tv)

	tv2 := TV{
		Goods{
			Price : 6999,
			Name : "电视机002",
		},
		Brand{
			Name : "创维",
			Address : "广州",
		},
	}

	fmt.Println("tv2",tv2)


	var child Child
	child.Name = "Bruce"
	child.Age = 22
	child.Gender = "male"
	child.Salary = 50000.00
	
	fmt.Println("child",child)
}