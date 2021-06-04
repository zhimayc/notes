package main

import (
	"fmt"
	"reflect"
)



func reflectTest(b interface{}){

	rVal := reflect.ValueOf(b)
	// Elem() 返回b持有的指针指向的值
	rVal.Elem().SetInt(20)

}

func main(){

	var num int = 10
	reflectTest(&num)
	fmt.Println(num)

}