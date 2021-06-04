package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}


//循环判断传入的参数类型
func TypeJudge(items ...interface{}) {
	for i,x := range items {
		switch x.(type) {	//这里type是一个关键字，固定写法
			case bool :
				fmt.Printf("param #%d is a bool 值是%v \n",i,x)
			case float32 :
				fmt.Printf("param #%d is a float32 值是%v \n",i,x)
			case float64 :
				fmt.Printf("param #%d is a float64 值是%v \n",i,x)
			case int,int32,int64 :
				fmt.Printf("param #%d is an int 值是%v \n",i,x)
			case nil :
				fmt.Printf("param #%d is nil 值是%v \n",i,x)
			case string :
				fmt.Printf("param #%d is a string 值是%v \n",i,x)
			case Point :
				fmt.Printf("param #%d is a Point 值是%v \n",i,x)
			case *Point :
				fmt.Printf("param #%d is a *Point 值是%v \n",i,x)
			default :
				fmt.Printf("param #%d's type is unknown 值是%v \n",i,x)
		}
	}
}

func main() {
	var a interface{}
	var point Point = Point{1,2}
	a = point

	var b Point
	b = a.(Point)		//类型断言
	fmt.Println(b)


	//带检测的类型断言
	var x interface{}
	var c float32 = 1.1
	x = c	//空接口可以接收任何类型

	if 	y,flag := x.(float32);flag {
		fmt.Println("convert successful")
		fmt.Printf("y的值类型是%T 值是=%v \n",y,y)
	} else {
		fmt.Println("convert failed")
	}

	fmt.Println("continue")


	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "Mike"
	address := "beijing"
	n4 := 300
	TypeJudge(n1,n2,n3,name,address,n4,b,&b)
	
}