package main

import (
	"fmt"
	"net"
)

func Process(conn net.Conn){

	defer conn.Close()	// 关闭连接
	for {
		buf := make([]byte,1024)

		// 1. 等待客户端通过conn发送信息
		// 2. 如果客户端没有Write(发送)，那么协程就阻塞在这里
		// fmt.Println("等待客户端输入中...")
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器的Read err = ",err)
			return
		}

		fmt.Print(string(buf[:n]))
	}

}


func main(){

	fmt.Println("服务器开始监听...")

	listen,err := net.Listen("tcp","0.0.0.0:8888")
	
	if err != nil {
		fmt.Println("listen err = ",err)
		return
	}

	defer listen.Close()   // 延时关闭listen

	// 循环等待客户端连接
	for {
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err = ",err)
		}else {
			fmt.Printf("Accept() suc conn = %v 客户端ip = %v \n",conn,conn.RemoteAddr().String())
		}

		// 准备一个协程为客户端服务
		go Process(conn)
	}


}