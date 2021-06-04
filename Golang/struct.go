package main

import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string `json:"name"`
	Age int	`json:"age"`
	Skill string `json:"skill"`
}

func main() {

	//1. 创建一个monster变量
	monster := Monster{"牛魔王",500,"芭蕉扇"}

	//2. 将monster变量序列化为json格式字符串
	jsonStr,err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 处理错误",err)
	}

	fmt.Println("jsonStr :",string(jsonStr))
}