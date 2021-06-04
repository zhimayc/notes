package main

import (
	"fmt"
	_"os"
	"flag"
)

func main() {


	// fmt.Println("命令行的参数有",len(os.Args))

	// for i,v := range os.Args {
	// 	fmt.Printf("args[%v] = %v \n",i,v)
	// }





	//定义几个变量，用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	//&user 就是接收用户命令行中输入的 -u 后面的参数值
	flag.StringVar(&user,"u","","用户名，默认为空")
	flag.StringVar(&pwd,"pwd","","密码，默认为空")
	flag.StringVar(&host,"h","localhost","主机，默认为空")
	flag.IntVar(&port,"p",3306,"端口，默认为3306")

	//转换，必须调用该方法
	flag.Parse()

	//输出结果
	fmt.Printf("user = %v pwd = %v host = %v port = %v \n",user,pwd,host,port)

}