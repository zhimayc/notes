package main

import (
	"fmt"
)

func main(){

	const (
		a = 1
		b = 2
	)

	const (
		c = iota	// c  = 0
		d		// d = c + 1
		e		// e = d + 1
	)

	// 仍然可以通过首字母的大小写来控制常量的访问范围

	fmt.Println("a = ",a)
	fmt.Println("b = ",b)


	fmt.Println("c = ",c)
	fmt.Println("d = ",d)
	fmt.Println("e = ",e)


}