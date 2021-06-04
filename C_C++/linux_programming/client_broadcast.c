#include <stdio.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>

#define SERV_PORT 8000
#define CLIENT_PORT 9000
#define MAXLINE 4096

int main(void){
    struct sockaddr_in localaddr;
    int confd;
    ssize_t len;
    char buf[MAXLINE];

    confd = socket(AF_INET,SOCK_DGRAM,0);

    bzero(&localaddr,sizeof(localaddr));
    localaddr.sin_family = AF_INET;
    inet_pton(AF_INET,"0.0.0.0",&localaddr.sin_addr.s_addr);
    localaddr.sin_port = htons(CLIENT_PORT);

    int ret = bind(confd,(struct sockaddr *)&localaddr,sizeof(localaddr));
    if(ret == 0){
        printf("...bind ok...\n");
    }

    while(1){
        len = recvfrom(confd,buf,sizeof(buf),0,NULL,0);
        write(STDOUT_FILENO,buf,len);
    }

    close(confd);

    return 0;
}
