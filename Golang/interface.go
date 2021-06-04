package main

import (
	"fmt"
)

//声明一个接口
type Usb interface {
	Start()
	Stop()
}

type Phone struct {

}

//让Phone实现Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

type Camera struct {

}

//让Camera也实现Usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

//计算机
type Computer struct {

}

//编写一个方法Working
func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}


//golang中的接口不需要显式实现，只要一个变量含有接口类型中的所有方法，那么这个变量就实现这个接口


type AInterface interface {
	test01()
}

type BInterface interface {
	test02()
}

type CInterface interface {
	AInterface
	BInterface
	test03()
}


//如果需要实现CInterface，就需要将AInterface,BInterface的方法都实现
type Stu struct {

}

func (stu Stu) test01() {

}

func (stu Stu) test02() {
	
}


func (stu Stu) test03() {
	
}



func main() {

	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)
	

	//一个自定义类型只有实现了了某个接口，才能将该自定义类型的实例（变量）赋给接口类型
	//只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型
	//接口中不能有任何变量

	var c CInterface = Stu{}
	fmt.Println(c)


}