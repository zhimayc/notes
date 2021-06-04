#include <stdio.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <net/if.h>

#define SERV_PORT 8000
#define CLIENT_PORT 9000
#define GROUP "239.0.0.2"

int main(void){
    struct sockaddr_in localaddr;
    int confd;
    ssize_t len;
    char buf[BUFSIZ];

    struct ip_mreqn group;

    confd = socket(AF_INET,SOCK_DGRAM,0);

    bzero(&localaddr,sizeof(localaddr));
    localaddr.sin_family = AF_INET;
    inet_pton(AF_INET,"0.0.0.0",&localaddr.sin_addr.s_addr);
    localaddr.sin_port = htons(CLIENT_PORT);

    bind(confd,(struct sockaddr *)&localaddr,sizeof(localaddr));

    inet_pton(AF_INET,GROUP,&group.imr_multiaddr);  /* 设置组地址 */
    inet_pton(AF_INET,"0.0.0.0",&group.imr_address);    /* 本地任意ip添加到组播组 */
    group.imr_ifindex = if_nametoindex("eth0"); /* 给出网卡名，转换为对应编号:eth0 --> 编号 命令：ip ad */

    setsockopt(confd,IPPROTO_IP,IP_ADD_MEMBERSHIP,&group,sizeof(group));     /* 组播权限 */

    while(1){
        len = recvfrom(confd,buf,sizeof(buf),0,NULL,0);
        write(STDOUT_FILENO,buf,len);
    }

    close(confd);

    return 0;
}
