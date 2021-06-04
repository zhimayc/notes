package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	_"io/ioutil"
)

func CopyFile(dstFileName string,srcFileName string) (written int64,err error){
	srcFile,err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err = %v \n",err)
	}

	defer srcFile.Close()

	//通过srcFile,获取到Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstFile,err := os.OpenFile(dstFileName,os.O_WRONLY | os.O_CREATE,0666)
	if err != nil {
		fmt.Printf("open file err = %v \n",err)
		return
	}

	defer dstFile.Close()

	//通过dstFile，获取到Writer
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer,reader)

}

func main() {

	//打开文件
	// file,err := os.Open("./test.txt")
	// if err != nil {
	// 	fmt.Println("open file error = ",err)
	// }

	// // fmt.Printf("file = %v",file)

	// //当函数退出时，要及时关闭file
	// defer file.Close()

	// //创建一个 *Reader，是带缓冲的
	// /*
	// 	const {
	// 		defaultBufSize	= 4096  //默认的缓冲区为4096
	// 	}
	// */
	// reader := bufio.NewReader(file)
	// //循环读取文件内容
	// for {
	// 	str,err := reader.ReadString('\n')	//读到一个换行就结束
	// 	if err == io.EOF {	// io.EOF 表示文件的末尾
	// 		break
	// 	}

	// 	//输出内容
	// 	fmt.Print(str)
	// }

	// fmt.Println("文件读取结束...")



	//使用ioutil.ReadFile一次性将文件读取到位
	// file := "./test.txt"
	// content,err := ioutil.ReadFile(file)
	// if err != nil {
	// 	fmt.Printf("read file err = %v \n",err)
	// }
	// //把读取到的内容显示到终端
	// fmt.Printf("%v",string(content))	//[]byte
	// //因为我们没有显式地Open文件，因此也不需要显式地Close文件
	// //因为文件的Open和Close被封装到ReadFile方法里


	// filePath := "./test1.txt"
	// file,err := os.OpenFile(filePath,os.O_WRONLY | os.O_CREATE,0666)

	/*
		文件打开模式	

		const (
			O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
			O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
			O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
			O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
			O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
			O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
			O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
			O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
		)
	
	*/
	// if err != nil {
	// 	fmt.Printf("open file err = %v \n",err)
	// 	return
	// }

	// //及时关闭file句柄
	// defer file.Close()

	//准备写入内容
	// str := "hello,world! \n"
	//写入时，使用带缓存的 *Writer
	// writer := bufio.NewWriter(file)
	// for i := 0; i < 5;i++ {
	// 	writer.WriteString(str)
	// }

	//因为writer是带缓存的，因此在调用WriteString方法时，其实内容是先写入缓存的
	//所以需要调用Flush方法，将缓冲的数据真正写入到文件中，否则文件中会没有数据
	// writer.Flush()




	// file1Path := "test.txt"
	// file2Path := "test1.txt"
	// data,err := ioutil.ReadFile(file1Path)
	// if err != nil {
	// 	//说明读取文件有错误
	// 	fmt.Printf("read file err = %v \n",err)
	// 	return
	// }

	// err = ioutil.WriteFile(file2Path,data,0777)
	// if err != nil {
	// 	fmt.Printf("write file err = %v \n",err)
	// }


	srcFile := "./1.jpeg"
	dstFile := "./abc.jpeg"
	_,err := CopyFile(dstFile,srcFile)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Printf("拷贝错误 err = %v \n",err)
	}





}