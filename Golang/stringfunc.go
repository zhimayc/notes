package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){

	//(1)统计字符串的长度，按字节 len(str)
	//golang的编码统一为utf-8 (ascii的字符（字母和数字）占一个字节，汉字占用3个字节)
	str := "hello"
	fmt.Println("(1) str len = ",len(str))

	str2 := "hello,北京"
	//字符串遍历，同时处理有中文的问题
	r := []rune(str2)
	for i := 0; i < len(r);i++ {
		fmt.Printf("字符 = %c \n",r[i])
	}




	//(2)字符串转整数  func Atoi(s string) (i int, err error)
	n,err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误",err)
	} else {
		fmt.Println("(2) 转成的结果是：",n)
	}




	//(3)整数转字符串
	str3 := strconv.Itoa(12345)
	fmt.Printf("(3) str3 = %v,str3 type = %T \n",str3,str3)


	//(4)字符串转 []byte
	var bytes = []byte("hello go")
	fmt.Printf("(4) bytes = %v \n",bytes)


	//(5)[]byte 转字符串
	str4 := string([]byte{97,98,99})
	fmt.Printf("(5) str4 = %v \n",str4)


	//(6)10进制转2,8，16进制:str = strconv.FormatInt(123,2)，返回对应的字符串
	str5 := strconv.FormatInt(12,2)
	fmt.Printf("(6) 12对应的二进制是=%v \n",str5)
	str5 = strconv.FormatInt(12,16)
	fmt.Printf("(6) 12对应的十六进制是=%v \n",str5)



	//(7)查找子串是否在指定的字符串中 func Contains(s, substr string) bool
	b := strings.Contains("seafood","foo")
	fmt.Printf("(7) b = %v\n",b)


	//(8)统计一个字符串中有几个指定的子串 func Count(s, sep string) int
	num := strings.Count("chinese","e")
	fmt.Printf("(8) num= %v\n",num)


	//(9)不区分大小写的字符串比较(==是区分字母大小写的)
	b = strings.EqualFold("abc","Abc")
	fmt.Printf("(9) b = %v\n",b)
	b = "abc" == "Abc"
	fmt.Printf("(9) b = %v\n",b)



	//(10)返回子串在字符串第一次出现的index值，如果没有返回-1  func Index(s, sep string) int
	index := strings.Index("NLT_abc","abc")
	fmt.Println("(10) index = ",index)



	//(11)返回子串在字符串中最后一次出现的index，如没有返回-1  func LastIndex(s, sep string) int
	index = strings.LastIndex("go golang","go")
	fmt.Printf("(11) index = %v \n",index)



	//(12)将指定子串替换成另外一个子串 func Replace(s, old, new string, n int) string  n指定替换几个，如果为-1则全部替换
	str = strings.Replace("go go hello","go","北京",-1)
	fmt.Printf("(12) str = %v \n",str)



	//(13) 按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组 func Split(s, sep string) []string
	strArr := strings.Split("hello,world,ok",",")
	for i := 0;i < len(strArr);i++ {
		fmt.Printf("str[%v] = %v \n",i,strArr[i])
	}
	fmt.Printf("(13) strArr = %v \n",strArr)


	//(14) 将字符串的字母进行大小写的转换  func ToLower(s string) string   func ToUpper(s string) string
	str = "goLang Hello"
	str = strings.ToLower(str)
	fmt.Printf("(14) str = %v\n",str)
	str = strings.ToUpper(str)
	fmt.Printf("(14) str = %v\n",str)


	//(15) 将字符串两边的空格去掉  func TrimSpace(s string) string
	str = strings.TrimSpace("       have a good time        ")
	fmt.Printf("(15) str = %q\n",str)



	//(16) 将字符串左右两边指定的字符去掉  func Trim(s string, cutset string) string
	str = strings.Trim(" !hell!o !","!")
	fmt.Printf("(16) str = %q\n",str)

	//(17) 将字符串左边指定的字符去掉 func TrimLeft(s string, cutset string) string

	//(18) 将字符串右边指定的字符去掉 func TrimRight(s string, cutset string) string


	//(19) 判断字符串是否以指定的字符串开头 func HasPrefix(s, prefix string) bool
	b = strings.HasPrefix("ftp://192.168.10.1","ftp")
	fmt.Printf("(19) b = %v \n",b)


	//(20) 判断字符串是否以指定的字符串结尾  func HasSuffix(s, suffix string) bool
	b = strings.HasSuffix("123abc.jpg","jpg")
	fmt.Printf("(20) b = %v \n",b)

	

}