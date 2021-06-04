package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"score"`
	Sex   string  `json:"sex"`	  
}

func (s Monster) Print() {
	fmt.Println("-----------start-------------")
	fmt.Println(s)
	fmt.Println("------------end--------------")
}

func (s Monster) GetSum(n1,n2 int) int {
	return n1 + n2;
}

func (s Monster) Set(name string,age int,score float32,sex string){
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(m interface{}){
	// 获取reflect.Type 类型
	typ := reflect.TypeOf(m)

	// 获取reflect.Value 类型
	val := reflect.ValueOf(m)

	// 获取m对应的类别
	kd  := val.Kind()

	// 如果传入的不是struct就退出
	if kd != reflect.Struct{
		fmt.Println("expect struct")
		return
	}

	// 获取结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d fields \n",num)

	// 遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d : 值为 = %v \n",i,val.Field(i))
		// 获取到struct标签，注意需要通过reflect.Type来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		// 如果该字段有tag标签就显示，否则就不显示
		if tagVal != "" {
			fmt.Printf("Field %d : tag为 = %v \n ",i,tagVal)
		}
	}

	// 获取该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods \n",numOfMethod)

	// var params [] reflect.Value
	// 反射调用方法是根据方法名的ASCII码来排序的，所以第一个方法是 GetSum()  第二个方法是Print()  第三个方法是Set()
	val.Method(1).Call(nil)  // 获取到第二个方法

	// 调用结构体的第1个方法 Method(0)
	var params []reflect.Value
	params = append(params,reflect.ValueOf(10))
	params = append(params,reflect.ValueOf(40))
	res := val.Method(0).Call(params)	// 传入的参数是 []reflect.Value
	fmt.Println("res=",res[0].Int())	// 返回结果，返回的结果是 []reflect.Value

}



func main(){

	var m Monster = Monster{
		Name  : "金角大王",
		Age   : 400,
		Score : 80.5,
	}

	TestStruct(m)

}