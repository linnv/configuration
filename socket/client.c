#include "include.h"
int checkConnection(int socket);
void connecting(int sock,struct sockaddr_in servaddr);
int main(int argc, char *argv[])
{
	int connfd,socketd;
	struct sockaddr_in servaddr;

	/*	if (argc !=2)
		{
		printf("usgae: tcpcli <IPaddress>");
		exit(1);

		}
		*/
	socketd = socket(AF_INET, SOCK_STREAM,0);   //sock_stream: tcp
	bzero(&servaddr, sizeof(servaddr));
	servaddr.sin_family = AF_INET;   //ipv4

	servaddr.sin_port = htons(3333);  //port: this must be the same with server
	servaddr.sin_addr.s_addr = inet_addr("127.0.0.9");  //ip address

	//inet_pton(AF_INET,argv[1],&servaddr.sin_addr);

	connecting(socketd,servaddr);

	//	if (connfd == 0)
	//	printf("connect success\n");	
	int max = 100;
	char buff[max];
	int n,i;
	for (i = 0; i < 1; ++i)
	{
	  int time,memory,statu;
//  		fscanf(socketd,"%d %d %d",&time,&memory,&statu);	  
				//printf("client: time:%d memory:%d statu:%d\n",time,memory,statu );
//		if(checkConnection(socketd))
		{

			n = read(socketd,buff,max);
			if (n ==0){
				printf("%d time read nothing \n",i);
			}
			else{
				sscanf(buff,"%d%d%d",&time,&memory,&statu);
				printf("client: time:%d memory:%d statu:%d\n",time,memory,statu );
				//printf("from server:%s",buff );	
			}

		}
//		sleep(1);
	}

	/*

	   if(checkConnection(socketd))
	   {
	   strcpy(buff,"calling server !\n");
	   write(socketd,buff,sizeof(buff));
	// sleep(3);
	n = read(socketd,buff,max);
	if (n ==0){
	printf("nothing to read\n");
	}
	else{
	printf("from server:%s",buff );	
	}

	//sleep(10);
	if(!checkConnection(socketd)){
	close(socketd);
	printf("reconnection\n");
//	connfd = connect(socketd,(struct sockaddr*)&servaddr,sizeof(servaddr));
connecting(socketd,servaddr);
checkConnection(socketd);
}
}
//	checkConnection(socketd);
//	*/

//exit(0);
return 0;
}
/*
 * return 1:if connect successfully
 */
int checkConnection(int socket)
{	
	if (socket <= 0)
	{
		return 0;

	}	
	struct tcp_info info;
	int length = sizeof(info);

	getsockopt(socket,IPPROTO_TCP,TCP_INFO,&info,(socklen_t*)&length);

	if(info.tcpi_state == 1)
	{	
		printf("connected!\n");	
		return 1;
	}
	else{
		printf("disconnect!\n");
		return 0;	
	}

}
void connecting(int sock,struct sockaddr_in servaddr)
{
	int n=-1;
	while(0!=n){
		/*
		 * The normal return value from connect is 0. If an error occurs, connect returns -1.
		 */
		n=(connect(sock,(struct sockaddr*)&servaddr,sizeof(servaddr)));
		printf("connecting\n");
	}

}
