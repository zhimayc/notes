package main

import(
   "fmt"
)

func main() {

	var intArr [5]int = [...]int{1,22,33,66,99}

	//引用intArr数组的起始下标为1，最后的下标为3（但是不包含3）
	slice := intArr[1:3]
	fmt.Println("intArr = ",intArr)
	fmt.Println("slice 的元素是 = ",slice)
	fmt.Println("slice 的元素个数 = ",len(slice))
	fmt.Println("slice 的容量 = ",cap(slice))	//切片的容量是可以动态变化的


	//切片slice是对数组intArr的引用
	fmt.Printf("intArr[1] 地址 = %p \n",&intArr[1])
	fmt.Printf("slice[0]地址 = %p \n",&slice[0])
	slice[0] = 55
	fmt.Println(intArr)

	//使用make 创建切片 通过make方式创建切片可以指定切片的大小和容量
	var slice1 []float64 = make([]float64,5,10)
	slice1[1] = 10
	slice1[3] = 20
	fmt.Println(slice1)

	//定义切片，直接就指定具体数组
	var slice2 []string = []string{"nginx","tomcat","apache"}
	fmt.Println(slice2)

	//引用数组全部元素简写
	slice3 := intArr[:]
	fmt.Println(slice3)

	//切片可以再切片
	slice4 := slice3[1:2]
	fmt.Println(slice4)


	//用append内置函数，可以对切片进行动态追加
	var slice5 []int = []int{100,200,300}
	//通过append直接给slice5增加具体的元素
	slice5 = append(slice5,400,500,600)
	fmt.Println(slice5)


	//切片的拷贝操作 使用内置函数copy完成拷贝  func copy(dst, src []Type) int
	var slice6 []int = []int {1,2,3,4,5}
	var slice7 = make([]int,10)
	copy(slice7,slice6)
	fmt.Println(slice7)


	//string和slice
	str := "hello,everybody"
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str = ",str)




}