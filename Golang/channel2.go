package main

import (
	"fmt"
)


//write Data
func writeData(intChan chan int) {
	for i := 0;i < 50;i++ {
		// 放入数据
		intChan<- i
		fmt.Println("writeData ",i)
	}
	close(intChan)
}

//read data
func readData(intChan chan int,exitChan chan bool) {
	for {
		v,ok := <-intChan
		if !ok {
			break;
		}

		fmt.Printf("readData 读到数据=%v\n",v)
	}
	exitChan<- true
	close(exitChan)
}

func main(){

	//创建两个管道
	intChan := make(chan int,20)
	exitChan := make(chan bool,1)

	go writeData(intChan)
	go readData(intChan,exitChan)

	for {
		_,ok := <-exitChan
		if !ok {
			break
		}
	}



	//管道可以声明为只读或者只写
	//1.在默认情况下管道是双向的

	// var chan1 chan int  // 可读可写

	//2. 声明为只写
	// var chan2 chan<- int 
	// chan2 = make(chan int,3)
	// chan2<- 20
	// fmt.Println("chan2 = ",chan2)

	//3. 声明为只读
	// var chan3 <-chan int 
	// num2 := <-chan3
	// fmt.Println("num2 = ",num2)





}