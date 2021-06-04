package main

import (
	"fmt"
)

func main(){

	//四种初始化数组的方式
	var arr[3] int = [3] int {1,2,5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	var arr2 = [3] int {1,2,9}
	fmt.Println("arr2 = ",arr2)

	var arr3 = [...] int {1,3,5,7,9}
	fmt.Println("arr3 = ",arr3)

	var arr4 = [...] int {1:800,0:900,2:999}
	fmt.Println("arr4 = ",arr4)

	var strarr = [...] string {1:"tom",0:"jack",2:"jobs"}
	fmt.Println("strarr = ",strarr)
}