package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"	// 引入redis包
)

func main(){

	// 通过go向redis写入数据和读取数据
	conn,err := redis.Dial("tcp","192.168.0.199:6379")

	if err != nil {
		fmt.Println("redis.Dial err = ",err)
		return
	}

	defer conn.Close()

	// 2. 通过go向redis写入数据 string [key-val]
	_,err = conn.Do("Set","name","tom lucy")

	if err != nil {
		fmt.Println("set err = ",err)
		return
	}

	fmt.Println("操作ok")


	//3. 通过go 向redis读取数据string [key-val]
	r,err := redis.String(conn.Do("Get","name"))
	if err != nil {
		fmt.Println("Get err = ",err)
		return
	}

	// 因为返回r是interface{}
	// 因为name对应的值是string,因此我们需要转换

	fmt.Println("name = ",r)




	// 操作Hash类型
	_,err = conn.Do("HSet","user01","name","Micheal")
	if err != nil {
		fmt.Println("hset err = ",err)
		return
	}

	_,err = conn.Do("HSet","user01","age",18)
	if err != nil {
		fmt.Println("hset err = ",err)
		return
	}

	r1,err := redis.String(conn.Do("HGet","user01","name"))
	if err != nil {
		fmt.Println("HGet err = ",err)
		return
	}

	fmt.Println("user01 name = ",r1)

	r2,err := redis.Int(conn.Do("HGet","user01","age"))
	if err != nil {
		fmt.Println("HGet err = ",err)
		return
	}

	fmt.Println("user01 age = ",r2)



	_,err = conn.Do("hmset","user02","name","Taylor Swift","age",20)
	if err != nil {
		fmt.Println("hmset err = ",err)
		return
	}

	r3,err := redis.Strings(conn.Do("hmget","user02","name","age"))
	if err != nil {
		fmt.Println("hmget err = ",err)
		return
	}

	for i,v := range r3 {
		fmt.Printf("r3[%d] = %s \n",i,v)
	}


	r4,err := redis.StringMap(conn.Do("hgetall","user02"))
	if err != nil {
		fmt.Println("hgetall err = ",err)
		return
	}

	for i,v := range r4 {
		fmt.Printf("r4 %s = %s \n",i,v)
	}




}