#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <string.h>
#include <arpa/inet.h>
#include <net/if.h>

#define SERV_PORT 8000
#define CLIENT_PORT 9000
#define MAXLINE 1500

#define BROADCAST_IP "192.168.1.255"

int main(int argc,char *argv[]){
    int sockfd;
    struct sockaddr_in serv_addr,clie_addr;
    char buf[MAXLINE];

    sockfd = socket(AF_INET,SOCK_DGRAM,0);

    bzero(&serv_addr,sizeof(serv_addr));
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_addr.s_addr = htonl(INADDR_ANY);
    serv_addr.sin_port = htons(SERV_PORT);
    bind(sockfd,(struct sockaddr *)&serv_addr,sizeof(serv_addr));

    int flag = 1;
    setsockopt(sockfd,SOL_SOCKET,SO_BROADCAST,&flag,sizeof(flag));

    bzero(&clie_addr,sizeof(clie_addr));
    clie_addr.sin_family = AF_INET;
    inet_pton(AF_INET,BROADCAST_IP,&clie_addr.sin_addr.s_addr);
    clie_addr.sin_port = htons(CLIENT_PORT);
    
    int i = 0;
    while(1){
        sprintf(buf,"Drink %d glasses of water \n",i++);
        sendto(sockfd,buf,strlen(buf),0,(struct sockaddr *)&clie_addr,sizeof(clie_addr));
        sleep(1);
    }

    close(sockfd);


    return 0;
}