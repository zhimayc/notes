package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"	// 引入redis包
)

// 定义一个全局的redis连接池 pool
var pool *redis.Pool

// 当启动程序时，就初始化连接池
func init(){
	pool = &redis.Pool {
		MaxIdle : 8,	    // 最大空闲连接数
		MaxActive : 0,      // 表示和数据库最大连接数，0表示没有限制
		IdleTimeout : 100,	// 最大空闲时间
		Dial : func() (redis.Conn,error) {	// 初始化连接
			return redis.Dial("tcp","192.168.0.199:6379")
		},
	}
}

func main(){

	// 先从pool中取出一个连接
	conn := pool.Get()
	defer conn.Close()

	_,err := conn.Do("Set","name","Justin billboard")
	
	if err != nil {
		fmt.Println("conn.Do err = ",err)
		return
	}

	r,err := redis.String(conn.Do("Get","name"))
	fmt.Printf("r = %v \n",r)

}