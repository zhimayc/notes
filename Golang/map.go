package main

import (
	"fmt"
	"sort"
)

func main() {
	//map的声明和注意事项
	var a map[string]string
	//在使用map前，需要先make，make的作用就是给map分配数据空间

	//第一种方式
	a = make(map[string]string,10)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no1"] = "武松"
	a["no3"] = "吴用"
	a["no4"] = "鲁智深"
	a["no5"] = "史进"
	fmt.Println(a)

	//删除map中的元素  func delete(m map[Type]Type1, key Type)
	//内建函数delete按照指定的键将元素从映射中删除。若m为nil或无此元素，delete不进行操作。
	delete(a,"no1")
	fmt.Println(a)


	//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "广州"
	fmt.Println(cities)

	//map的查找
	val,ok := cities["no1"]
	if ok {
		fmt.Printf("有no1 key 值为 %v \n",val)
	} else {
		fmt.Printf("没有no1 key \n")
	}

	//map的遍历
	for k,v := range cities {
		fmt.Printf("k = %v v = %v \n",k,v)
	}

	//map的长度
	fmt.Println("cities len = ",len(cities))


	//第三种方式
	heroes := map[string]string{
		"hero1" : "李逵",
		"hero2" : "鲁智深",
		"hero3" : "林冲",
	}

	fmt.Println("heroes = ",heroes)

	students := map[string]map[string]string{
		"1101" : {
			"name":"lucy",
			"gender":"female",
		},
		"1102" : {
			"name":"jane",
			"gender":"female",
		},
	}

	fmt.Println(students)

	studentsMap := make(map[string]map[string]string)
	studentsMap["stu01"] = make(map[string]string,2)
	studentsMap["stu01"]["name"] = "Alice"
	studentsMap["stu01"]["gender"] = "female"

	studentsMap["stu02"] = make(map[string]string,2)
	studentsMap["stu02"]["name"] = "Zed"
	studentsMap["stu02"]["gender"] = "male"

	fmt.Println(studentsMap)


	//map切片的使用
	var monsters []map[string]string
	monsters = make([]map[string]string,2)

	if monsters[0] == nil {
		monsters[0] = make(map[string]string,2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string,2)
		monsters[1]["name"] = "玉兔精"
		monsters[1]["age"] = "400"
	}

	//使用切片的append函数，可以动态的增加monster
	newMonster := map[string]string {
		"name" : "金角大王~火云邪神",
		"age" : "200",
	}

	monsters = append(monsters,newMonster)

	fmt.Println(monsters)


	//map的排序
	map1 := make(map[int]int,10)
	map1[10] = 100
	map1[1]  = 13
	map1[4]  = 56
	map1[8]  = 90

	fmt.Println(map1)

	//如果按照map的key的顺序进行排序输出
	//1. 先将map的key放入到切片中
	//2. 对切片排序
	//3. 遍历切片，然后按照key来输出map的值

	var keys []int
	for k,_ := range map1 {
		keys = append(keys,k)
	}

	//排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _,k := range keys {

		fmt.Printf("map1[%v] = %v \n",k,map1[k])
	}

	//map是引用类型,遵循引用类型传递的机制

	type Stu struct {
		Name string
		Age int
		Address string
	}

	students2 := make(map[string]Stu,10)
	stu1 := Stu{"tom",18,"New York"}
	stu2 := Stu{"mary",16,"London"}
	students2["no1"] = stu1
	students2["no2"] = stu2
	fmt.Println(students2)


	//遍历各个学生信息
	for k,v := range students2 {
		fmt.Printf("学生的编号是 %v \n",k)
		fmt.Printf("学生的名字是 %v \n",v.Name)
		fmt.Printf("学生的年龄是 %v \n",v.Age)
		fmt.Printf("学生的地址是 %v \n",v.Address)
		fmt.Println()
	}


}