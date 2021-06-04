package main

import (
	_"fmt"
	"testing"  //testing 提供对 Go 包的自动化测试的支持
)

func TestSum(t *testing.T){
	t.Logf("Sum 执行正确！")
}

//编写测试用例，测试addUpper是否正确
//func TestXxx(*testing.T) 其中 Xxx 可以是任何字母数字字符串（但第一个字母不能是 [a-z]），用于识别测试例程。
func TestAddUpper(t *testing.T) {
	//调用
	res := addUpper(10)
	if res != 55 {
		// fmt.Printf("addUpper(10) 执行错误！期望值 = %v,实际值 = %v \n",55,res)
		t.Fatalf("addUpper(10) 执行错误！期望值 = %v,实际值 = %v \n",55,res)
	}

	//如果正确，输出日志
	t.Logf("addUpper(10) 执行正确！")

}


/**

	注意事项：
	1) 测试用例文件名必须以 _test.go结尾。比如cal_test.go,cal不是固定的。
	2) 测试用例函数必须以Test开头，一般来说就是Test+被测试的函数名，比如TestAddUpper
	3) TestAddUpper(t *testing.T) 的形参类型必须是 *testing.T
	4) 一个测试用例文件中，可以有多个测试用例函数，比如TestAddUpper、TestSub
	5) 运行测试用例指令
	  a cmd>go test [如果运行正确，无日志，错误时，会输出日志]
	  b cmd>go test -v [运行正确或是错误，都输出日志]
	6) 当出现错误时，可以使用t.Fatalf来格式化输出错误信息，并退出程序
	7) t.Logf方法可以输出相应的日志
	8) 测试用例函数，并没有放在main函数中，也执行了，这就是测试用例的方便之处
	9) PASS 表示测试用例运行成功，FAIL 表示测试用例运行失败
	10) 测试单个文件，一定要带上被测试的原文件
	    go test -v cal_test.go cal.go
	11) 测试单个方法
	    go test -v -run TestAddUpper



*/
