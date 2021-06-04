package main

import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string `json:"name"` //tag标签
	Age int `json:"age"`
	Birthday string `json:"birth"`
	Sal float64	`json:"sal"`
	Skill string `json:"skill"`
}

//将结构体进行序列化
func testStruct() {
	monster := Monster{
		Name : "牛魔王",
		Age : 500,
		Birthday : "1900-11-11",
		Sal : 8000.0,
		Skill : "芭蕉扇",
	}

	data,err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化失败 err = %v \n",err)
	}

	//输出序列化后的结果
	fmt.Printf("struct 序列化后的结果 = %v \n",string(data))
}


//将map进行序列化
func testMap() {
	//定义一个map
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "火云洞"

	//将a这个map进行序列化
	data,err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化失败 err = %v \n",err)
	}

	//输出序列化后的结果
	fmt.Printf("Map 序列化后的结果 = %v \n",string(data))

}

//演示对切片进行序列化
func testSlice(){

	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 7
	m1["address"] = "London"
	slice = append(slice,m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 16
	m2["address"] = [2]string{"Tokyo","Beijing"}
	slice = append(slice,m2)

	//将切片进行序列化操作
	data,err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化失败 err = %v \n",err)
	}

	//输出序列化后的结果
	fmt.Printf("slice 序列化后的结果 = %v \n",string(data))

}


//将json字符串反序列化成struct
func unmarshalStruct() {
	str := "{\"name\":\"牛魔王\",\"age\":500,\"birth\":\"1900-11-11\",\"sal\":8000,\"skill\":\"芭蕉扇\"}"
	var monster Monster
	err := json.Unmarshal([]byte(str),&monster)
	if err != nil {
		fmt.Printf("unmarshal err = %v \n",err)
	}

	fmt.Printf("反序列化后 struct = %v  monster.Skill = %v \n",monster,monster.Skill)
}

//将json字符串反序列化成map
func unmarshalMap() {
	str := "{\"address\":\"火云洞\",\"age\":30,\"name\":\"红孩儿\"}"
	var a map[string]interface{}
	err := json.Unmarshal([]byte(str),&a)
	if err != nil {
		fmt.Printf("unmarshal err = %v \n",err)
	}

	fmt.Printf("反序列化后 map = %v \n",a)
}

//将json字符串反序列化成切片
func unmarshalSlice() {
	str := "[{\"address\":\"London\",\"age\":7,\"name\":\"jack\"},{\"address\":[\"Tokyo\",\"Beijing\"],\"age\":16,\"name\":\"tom\"}]"
	var slice []map[string]interface{}
	//反序列化不需要make，因为make操作被封装到Unmarshal方法里
	err := json.Unmarshal([]byte(str),&slice)
	if err != nil {
		fmt.Printf("unmarshal err = %v \n",err)
	}

	fmt.Printf("反序列化后 slice = %v \n",slice)

}

func main() {

	//演示将结构体,map，切片进行序列化

	//结构体进行序列化
	testStruct()

	//map进行序列化
	testMap()

	//切片进行序列化
	testSlice()

	//json反序列化成结构体
	unmarshalStruct()

	//json反序列化成map
	unmarshalMap()

	//json反序列化成切片
	unmarshalSlice()
	

}
