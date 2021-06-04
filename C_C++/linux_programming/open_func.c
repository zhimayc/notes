#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>


int main(){
    int fd;

    //打开已经存在的文件
    // fd = open("1.txt",O_RDWR);   //以只读方式打开
    // if(fd == -1){
    //     perror("open file error");
    //     exit(1);
    // }

    //创建新文件

    // fd = open("myhello",O_RDWR | O_CREAT,0777);
    // if(fd == -1){
    //     perror("open file error");
    //     exit(1);
    // }

    // printf("fd = %d \n",fd);


    //判断文件是否存在
    // fd = open("myhello",O_RDWR | O_CREAT | O_EXCL,0777);
    // if(fd == -1){
    //     perror("open file error");
    //     exit(1);
    // }



    //将文件截断为0
    fd = open("myhello",O_RDWR | O_TRUNC);
    if(fd == -1){
        perror("open file error");
        exit(1);
    }







    //关闭文件
    int ret = close(fd);
    printf("ret = %d \n",ret);




    return 0;
}












