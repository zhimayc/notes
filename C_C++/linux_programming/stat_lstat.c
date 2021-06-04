#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>

/*

stat 函数，是穿透函数，当读软链接时，读的是软链接指向的文件的信息
lstat 函数，是非穿透的，当读软链接时，读的是软链接本身的信息

*/

int main(int argc,char * argv[]){

    if(argc < 2){
        printf("./a.out filename \n");
        exit(1);
    }

    struct stat st;

    int ret = stat(argv[1],&st);
    if(ret == -1){
        perror("stat");
        exit(1);
    }

    //获取文件大小
    int size = (int)st.st_size;
    printf("file size = %d\n",size);

    return 0;
}












