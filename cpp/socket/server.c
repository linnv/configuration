#include "include.h"
int checkConnection(int socket);
int main(int argc, char *argv[])

{
	int listend, connfd;
	socklen_t len;
	//	struct sockaddr servaddr, clienAddress;
	struct sockaddr_in servaddr, clienAddress;
	char buff[100];
	time_t ticks;

	/*
	 * create the server's socket
	 */
	listend = socket(AF_INET,SOCK_STREAM,0);
	/*
	 * set addr for bind();
	 */
	bzero(&servaddr,sizeof(servaddr));
	servaddr.sin_family = AF_INET;

	//	servaddr.sin_addr.s_addr = inet_addr("1.1.1.1");
	servaddr.sin_addr.s_addr = htonl(INADDR_ANY);	
	servaddr.sin_port = htons(3333);

	/*
	 * bind the ip and port etc to the server's socket
	 */
	bind(listend,&servaddr,sizeof(servaddr));

	/*
	 * active the server's socket
	 */
	listen(listend,2);

	int count = 0 ;
	for(; ; ){
		//waiting for a client to connect to the server
		len = sizeof(clienAddress);
		connfd = accept(listend, (struct sockaddr*) & clienAddress, & len);	
		/*
		 * print the client's info getted by the accept()
		 *inet_ntop(): convertthe 32-bit IP address in the socket address structure to a dotted-decimal ASCII string
		 *ntohs(): convert the 16-bit port number from network byte order to host byte order.
		 */
		printf("connection from %s, port: %d\n",
				inet_ntop(AF_INET, & clienAddress.sin_addr,buff,sizeof(buff)),
				ntohs(clienAddress.sin_port));
		{
			int i;
			for (i = 0; i < 1; ++i)
			{
				/*
				 * get the local time
				 */
				ticks = time(NULL);
				/*
				 * using ctime format the local time getted and put it to buff
				 */
				/*
				snprintf(buff,sizeof(buff),"new %.24s\r\n",ctime(&ticks));
				printf("system time: %s\n",buff );
				 * sent to the client connected
				 */
				if(checkConnection(connfd)){

					snprintf(buff,sizeof(buff),"%d %d %d\n",1,2,3);
					write(connfd,buff,strlen(buff));	
					printf("second write done\n");


					//	read(connfd,buff,sizeof(buff));
					//while( read(connfd,buff,sizeof(buff)))
					//	printf("info from clien: \n",buff);

				}
//				sleep(1);
			}
		}

		close(connfd);

		return 0;
	}
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


	printf("disconnect!\n");
	return 0;	

}
