package main

import (
	"fmt"
)

type Person struct {
	Name string
}

//给Person结构体添加speak方法
func (p Person) speak() {
	fmt.Println(p.Name,"is a good man")
}

func (p Person) calculate(n int) {
	res := 0
	for i := 0;i <= n;i++ {
		res += i
	}
	fmt.Println(p.Name,"计算的结果是 :",res)
}

func (p Person) getSum(n1 int,n2 int) int {

	return n1 + n2
}

//给Person类型绑定方法
func (p Person) test() {
	fmt.Println("test() ",p.Name)
}

func main(){
	var p Person
	p.Name = "tom"
	p.test()	//调用方法
	p.speak()
	p.calculate(100)
	res := p.getSum(10,20)
	fmt.Println("res = ",res)

}