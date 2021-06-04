#include <stdio.h>
#include <unistd.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <stdlib.h>
#include <ctype.h>

#define SERV_IP "127.0.0.1"
#define SERV_PORT 6666
int main(void){
    int lfd,cfd;
    struct sockaddr_in  serv_addr,clie_addr;
    socklen_t clie_addr_len;
    char buf[BUFSIZ],clie_IP[BUFSIZ];
    int n,i,ret;
    lfd = socket(AF_INET,SOCK_STREAM,0);
    if(lfd == -1){
        perror("socket error");
        exit(1);
    }
    serv_addr.sin_family = AF_INET; //应用的协议是IPV4协议
    serv_addr.sin_port = htons(SERV_PORT); //绑定端口号  需要把主机字节序转换成网络字节序
    serv_addr.sin_addr.s_addr = htonl(INADDR_ANY);  //绑定本地有效的ip，需要把主机字节序转换成网络字节序
    ret = bind(lfd,(struct sockaddr *)&serv_addr,sizeof(serv_addr));
    if(ret == -1){
        perror("bind error");
        exit(1);
    }

    ret = listen(lfd,128);  //最大允许128个连接

    if(ret == -1){
        perror("listen error");
        exit(1);
    }

    cfd = accept(lfd,(struct sockaddr *)&clie_addr,&clie_addr_len);
    if(cfd == -1){
        perror("accept error");
        exit(1);
    }

    printf("client ip: %s,client port: %d \n",inet_ntop(AF_INET,&clie_addr.sin_addr.s_addr,clie_IP,sizeof(clie_IP)),ntohs(clie_addr.sin_port));

    while(1){
         n = read(cfd,buf,sizeof(buf));
        for(i = 0;i < n;i++){
            buf[i]  = toupper(buf[i]);
        }
        write(cfd,buf,n);
    }
   

    close(lfd);
    close(cfd);
}