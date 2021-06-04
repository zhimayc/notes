#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>


int main(){

    //打开一个已经存在的文件
    int fd = open("1.txt",O_RDONLY);
    if(fd == -1){
        perror("open file error");
        exit(1);
    }

    //创建一个新文件--写操作
    int fd1 = open("newfile",O_CREAT | O_WRONLY,0664);
    if(fd == -1){
        perror("open newfile error");
        exit(1);
    }

    //read file
    char buf[2048] = {0};
    int count = read(fd,buf,sizeof(buf));  
    if(count == -1){
        perror("read");
        exit(1);
    }

    while(count){
        //将读出的数据写入到另一个文件中
        int ret = write(fd1,buf,count);
        printf("write bytes %d \n",ret);
        //继续读文件
        count = read(fd,buf,sizeof(buf)); 
    }

    //关闭文件
    close(fd);
    close(fd1);


    return 0;
}












