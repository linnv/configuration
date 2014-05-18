#include "include.h"
int checkConnection(int socket);
int main(int argc, char *argv[])

{
	int listend, connfd;
	socklen_t len;
	struct sockaddr_in servaddr, cliaddr;
	char buff[100];
	time_t ticks;

	//create the server's socket
	listend = socket(AF_INET,SOCK_STREAM,0);
	//set addr for bind();
	bzero(&servaddr,sizeof(servaddr));
	servaddr.sin_family = AF_INET;
//	servaddr.sin_addr.s_addr = inet_addr("1.1.1.1");	
	servaddr.sin_addr.s_addr = htonl(INADDR_ANY);	
	servaddr.sin_port = htons(3333);
	
	//bind the ip and port etc to the server's socket
	bind(listend,&servaddr,sizeof(servaddr));

	//active the server's socket
	listen(listend,2);
	
	int count = 0 ;
	for(; ; ){
		//waiting for a client to connect to the server
		len = sizeof(cliaddr);
	 connfd = accept(listend, (struct sockaddr*) & cliaddr, & len);	
	 //print the client's info getted by the accept()
	 //inet_ntop(): convertthe 32-bit IP address in the socket address structure to a dotted-decimal ASCII string
	 //ntohs(): convert the 16-bit port number from network byte order to host byte order.
	 //
	 printf("connection from %s, port: %d\n",
			 inet_ntop(AF_INET, & cliaddr.sin_addr,buff,sizeof(buff)),
			 ntohs(cliaddr.sin_port));
	 //get the local time
	 ticks = time(NULL);
	 //format the local time getted and put it to buff
	snprintf(buff,sizeof(buff),"new %.24s\r\n",ctime(&ticks));
	//sent to the client connected
	printf("system time: %s\n",buff );
	if(checkConnection(connfd)){
	//	sleep(2);
	write(connfd,buff,strlen(buff));	
	printf("write done\n");
	read(connfd,buff,sizeof(buff));
	//while( read(connfd,buff,sizeof(buff)))
		printf("info from clien: \n",buff);

	}
	//close the connection
	close(connfd);
	 connfd = accept(listend, (struct sockaddr*) & cliaddr, & len);	
//	sleep(2);
	count++;
	checkConnection(connfd);
	printf("\nconection done %d\n", count);
	
	}
	return 0;
}
int checkConnection(int socket){
	if (socket <= 0)
	{
		return 0;
	}	
	struct tcp_info info;
	int length = sizeof(info);
	getsockopt(socket,IPPROTO_TCP,TCP_INFO,&info,(socklen_t*)&length);
	if(info.tcpi_state == 1)
	{
		printf("check connected!\n");	
		return 1;
	}
	else{
	printf("disconnect!\n");
	return 0;	
	}

}
