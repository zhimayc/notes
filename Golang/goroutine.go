package main

import (
	"fmt"
	"time"
	"sync"
)

var (
	myMap = make(map[int]int,10)

	//声明一个全局的互斥锁
	lock sync.Mutex
)

// 编写一个函数，来计算各个数的累加结果，并放入到map中
// 启动多个协程，将统计的结果放入到map中
// map 应该做成一个全局的
func test(n int){
	res := 0
	for i := 1;i <= n;i++ {
		res += i
	}

	
	//加锁
	lock.Lock()

	//将res放入到myMap
	myMap[n] = res

	//解锁
	lock.Unlock()

}


func main(){

	//返回本地机器的逻辑CPU个数。
	// num := runtime.NumCPU()

	//设置可同时执行的最大CPU数
	// runtime.GOMAXPROCS(num)

	// fmt.Println("num = ",num)



	//开启200个协程，计算结果
	for i := 1; i <= 200;i++ {
		go test(i)
	}

	time.Sleep(time.Second * 5)
	//输出结果

	lock.Lock()
	for i,v := range myMap {
		fmt.Printf("map[%d] = %d \n",i,v)
	}
	lock.Unlock()

}