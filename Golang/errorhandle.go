package main

import (
	"fmt"
	"errors"
)

func test(){

	//使用defer + recover来捕获和处理异常
	defer func(){
		err := recover()  //recover()内置函数，可以捕获到异常
		if err != nil {
			fmt.Println("err = ",err)
		}
	}()
	num1 := 10
	num2 := 0
	res  := num1 / num2
	fmt.Println("res = ",res)
}


//函数去读取配置文件config.ini的信息
//如果文件名传入不正确，我们就返回一个自定义的错误
func readConf(name string)(err error){
	if name == "config.ini" {
		//正确
		return nil
	} else {
		//读取错误，返回一个自定义的错误
		return errors.New("读取文件错误...")
	}
}

func test02(){
	err := readConf("config2.ini")
	if err != nil {
		//如果读取文件发生错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("test02() 继续执行...")
}

func main(){
	// test()
	// fmt.Println("main() 继续执行...")

	//测试自定义错误的使用
	test02()
	fmt.Println("main() 继续执行...")


}