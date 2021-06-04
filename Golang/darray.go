package main

import(
	"fmt"
)

func main(){

	//定义或声明二维数组
	var arr[4][6]int
	//赋初值
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	//遍历二维数组
	for i := 0;i < 4;i++ {
		for j := 0;j < 6;j++ {
			fmt.Print(arr[i][j]," ")
		}
		fmt.Println()
	}

	var arr2 [2][3]int = [2][3]int {{1,2,3},{4,5,6}}
	fmt.Println(arr2)

	//for循环遍历二维数组
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]);j++ {
			fmt.Printf("%v\t",arr2[i][j])
		}
		fmt.Println()
	}

	//for-range遍历二维数组
	for i,v := range arr2 {

		for j,v2 := range v {
			fmt.Printf("arr2[%v][%v] = %v \t",i,j,v2)
		}

		fmt.Println()
	
	}
	
}