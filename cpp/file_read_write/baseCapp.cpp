#include <iostream>
#include "include.h"
using namespace std;
int Readn(int fd, void* buffer, size_t count) { 
     	char* p = (char*)buffer;
    while (count > 0 ) {   
       	    ssize_t num = read(fd, p, count);
            if (num == -1) { 
		    printf("Fail to read from file");
	                return -1;
	            }
            if (num == 0) {            // EOF
	                break;
	            }
            p += num;
            count -= num;
        }

    return p - (char*)buffer;
}

int Writen(int fd, const void* buffer, size_t count) 
{    const char*p = (const char*)buffer;
    while (count > 0 ) {      
	    int num = write(fd, p, count);
            if (num == -1) {    
		    printf("Fail to write from file");
	                return -1;
	            }
            p += num;
            count -= num;
        }
    return 0;
}
int main(int argc, char *argv[])
{
	int fd= open("./filetotest",O_RDWR|O_CREAT,S_IROTH|S_IRWXG|S_IRWXU);
	char buff[]="how a judge system work!\n";
	if(!Writen(fd,buff,sizeof(buff)))
			{
			printf("write done!\n");	
			}
	else{
	printf("fail to write\n");	
	}
	close(fd);
	memset(buff,0,sizeof(buff));
	printf("buff:%s sizeof(buff):%d\n",buff,sizeof(buff));
	char bf[100];
	int ff = open("./filetotest",O_RDWR);
	printf("bf's length%d\n", sizeof(ff));
	if(Readn(fd,bf,sizeof(bf)))
			{
				close(ff);
			printf("read done! %s\n",bf);	
			}
	else{
	printf("fail to read\n");	
	}
	



	return 0;
}
