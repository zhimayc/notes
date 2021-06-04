#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>
#include <fcntl.h>



int main(){

    int fd = open("tempfile",O_CREAT | O_RDWR,0664);
    if(fd == -1){
        perror("open");
        exit(1);
    }

    //删除临时文件
    int ret = unlink("tempfile");

    //write file
    write(fd,"hello\n",5);

    //重置文件指针
    lseek(fd,0,SEEK_SET);


    //read file
    char buf[24] = {0};
    int len = read(fd,buf,sizeof(buf));

    //将读出的内容，写到屏幕上
    write(1,buf,len);
    close(fd);

    return 0;
}












