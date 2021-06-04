package main

import (
	_"fmt"
)

//一个被测试函数
func addUpper(n int) int {
	res := 0
	for i := 1;i <= n;i++ {
		res += i
	}

	return res
}
