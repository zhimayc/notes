package main

import(
	"fmt"
	"time"
)

func main(){

	//(1) 获取当前时间
	now := time.Now()
	fmt.Printf("(1) now = %v,now type = %T \n",now,now)

	//(2) 通过now可以获取到年月日，时分秒
	fmt.Printf("(2) %v年%v月%v日 %v:%v:%v \n",now.Year(),int(now.Month()),now.Day(),now.Hour(),now.Minute(),now.Second())

	//(3) 格式化日期，时间
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Printf("\n")
	fmt.Printf(now.Format("15:04:05"))
	fmt.Printf("\n")

	//(4) 时间常量
	/*
		const (
			Nanosecond  Duration = 1
			Microsecond          = 1000 * Nanosecond
			Millisecond          = 1000 * Microsecond
			Second               = 1000 * Millisecond
			Minute               = 60 * Second
			Hour                 = 60 * Minute
		)
	
	*/
	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Second)
		if i > 5 {
			break
		}
	}


	//(5) Unix和UnixNano的使用
	fmt.Printf("unix时间戳 = %v unixnano时间戳 = %v \n",now.Unix(),now.UnixNano())



}

