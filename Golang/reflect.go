package main

import (
	"fmt"
	"reflect"
)



type Student struct {
	Name string
	Age int
	Gender string
}

func reflectTest(b interface{}){

	// 先获取reflect.Type 
	rType := reflect.TypeOf(b)
	fmt.Println("rType = ",rType)
	//获取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal = %v  rVal type = %T \n",rVal,rVal)

	n2 := 2 + rVal.Int()
	fmt.Printf("n2 = %v \n",n2)

	// 下面我们将 rVal 转成 interface{}
	iv := rVal.Interface()
	//将interface{} 通过断言转成需要的类型
	num2 := iv.(int)
	fmt.Println("num2 = ",num2)

}

func reflectTest2(b interface{}){
	// 先获取reflect.Type 
	rType := reflect.TypeOf(b)
	fmt.Println("rType = ",rType)
	//获取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal = %v  rVal type = %T \n",rVal,rVal)

	// 获取变量对应的kind
	// (1) rVal.Kind()
	// (2) rType.Kind()
	
	fmt.Printf("kind = %v kind = %v \n",rVal.Kind(),rType.Kind())


	// 下面我们将 rVal 转成 interface{}
	iv := rVal.Interface()
	fmt.Printf(" iv = %v iv type = %T \n",iv,iv)
	//将interface{} 通过断言转成需要的类型
	stu,ok := iv.(Student)
	if ok {
		fmt.Printf("stu.Name = %v \n",stu.Name)
	}
}

func main(){

	// 定义一个int
	var num int = 100
	reflectTest(num)

	// 定义一个Student的实例
	stu := Student{
		Name : "tom",
		Age : 20,
	}	

	reflectTest2(stu)



}