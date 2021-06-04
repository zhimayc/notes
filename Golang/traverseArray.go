package main

import (
	"fmt"
)

//修改原数组内容需要传指针
func test(arr *[3] int){
	(*arr)[0] = 88
}

func test1(arr [3] int){
	arr[0] = 99
}

func main(){

	var arr = [...] string{"lucy","jane","alice"}
	for _,value := range arr {
		fmt.Println("value = ",value)
	}

	var arr1[3] int = [3] int {1,2,3}
	test(&arr1)
	fmt.Println("arr1 = ",arr1)


	var arr2 = [...] int {4,5,6}
	test1(arr2)
	fmt.Println("arr2 = ",arr2)

	var myChar[26] byte
	for i := 0;i < 26;i++ {
		myChar[i] = 'A' + byte(i)
	}

	for i := 0; i < 26;i++ {
		fmt.Printf("%c ",myChar[i])
	}
}