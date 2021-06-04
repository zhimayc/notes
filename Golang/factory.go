package main

import (
	"fmt"
	"go_code/project1/model"
)

func main() {

	var stu = model.NewStudent("jack",90)
	fmt.Println(stu)
	fmt.Println("name = ",stu.Name," score = ",stu.Score)
}