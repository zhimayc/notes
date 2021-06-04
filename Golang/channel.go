package main

import (
	"fmt"
)


//write Data
func writeData(intChan chan int) {
	for i := 0;i < 50;i++ {
		// 放入数据
		intChan<- i
	}
	close(intChan)
}

//read data
func readData(intChan chan int,exitChan chan bool) {
	for {
		v,ok := <-intChan
		if !
	}
}

func main(){

	//创建两个管道
	intChan := make(chan int,50)
	exitChan := make(chan bool,1)

}