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

    int ret = lseek(fd,0,SEEK_END);
    printf("file length = %d\n",ret);

    //文件拓展
    ret = lseek(fd,2000,SEEK_END);
    printf("return value %d \n",ret);

    //实现文件拓展，需要在最后做一次写操作
    write(fd,"a",1);


    close(fd);

    return 0;
}












