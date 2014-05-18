//
//  main.cpp
//  traceClass
//
//  Created by wujialin on 11/3/14.
//  Copyright (c) 2014 wujialin. All rights reserved.
//

//fork
#include <unistd.h>
#include "functions.h"
#include <iostream>
using namespace std;
int main(int argc, const char * argv[])
{

	int pid = fork();
	if (pid <0)
	{
		cout<<"fork failed!"<<endl;
	}
	else if (pid == 0 )
	{
		execl("./app","./app"," ",(char * )0);
		cout<<"execute failed!"<<endl;
	}
	else{
		int status,sig;
		struct rusage  usage; 
		for(;;){
			updateUsage(pid);
			wait4(pid,&status,0,&usage);
			if (WIFEXITED(status))
			{
				sig = WEXITSTATUS(status);
				if (sig == 0)
				{
				printf("p info: app exit done!\n");
				}
				break;
			}
			if (WIFSIGNALED(status))
			{
				sig = WTERMSIG(status);
				switch(sig){
				case 0:
				printf("p info: app exit done!\n");
				break;
				case SIGALRM:
					alarm(0);	
				case SIGKILL:
				case SIGXCPU:
					printf("time exceed!\n");
					break;
				default:
					printf("other reason!\n");
					break;
				}
				break;
			}
		
		}
		int getTime = (usage.ru_stime.tv_sec + usage.ru_utime.tv_sec)*1000*1000 + (double)(usage.ru_stime.tv_usec + usage.ru_utime.tv_usec);
		//int getTime = (usage.ru_stime.tv_sec + usage.ru_utime.tv_sec)*1000 + (double)(usage.ru_stime.tv_usec + usage.ru_utime.tv_usec)/1000;
		printf("time getted by rusage:%d micro seconds\n",getTime );
		//printf("time by rusage:%d milliseconds\n",getTime );
			
	}
	
   return 0;

}

