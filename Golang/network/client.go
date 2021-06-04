package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)


func main(){

	conn,err := net.Dial("tcp","127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err = ",err)
		return
	}

	reader := bufio.NewReader(os.Stdin)	// os.Stdin 代表标准输入

	for {

		line,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err = ",err)
		}

		// 如果用户输入的是exit 就退出
		line = strings.Trim(line," \r\n")
		if line == "exit" {
			fmt.Println("客户端退出...")
			break
		}
	
		_,err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err = ",err)
		}
	}

}