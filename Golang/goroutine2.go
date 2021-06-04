package main

import (
	"fmt"
	"time"
)

func sayHello(){
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}

func test(){
	// 使用defer + recover 
	defer func(){
		// 捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误",err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang"	// error  没有先make 所以会报错
}

func main(){
	
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok ")
		time.Sleep(time.Second)
	}

}