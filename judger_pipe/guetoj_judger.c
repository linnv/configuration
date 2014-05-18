#include <unistd.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/time.h>
#include <sys/wait.h>
#include <sys/timeb.h>
#include <sys/syscall.h>
#include <sys/resource.h>
#include <sys/ptrace.h>
#include <sys/user.h>
#include <sys/reg.h>
#include <string.h>

#include "disabled_syscall.h"
#define LENGTH 65535
int timeLimitation;
int memoryLimitation;
struct rusage executableResourceUsage;
struct stRunStatus
{
    int timeUsed;
    int memoryUsed;
    int runExitStatus;
};

pid_t pid, subPid;     
int pi[2];
char memoryPath[LENGTH];
char buffer[LENGTH];
char resultFile[LENGTH];
char inputFile[LENGTH];
char outputFile[LENGTH];
long  timeConsumption;
long memoryConsumption;	

enum exitStatus {COMPILING = 100000, ACCEPTED, PRESENTATION_ERROR, TIME_LIMIT_ERROR, MEMORY_LIMIT_ERROR, WRONG_ANSWER, RUNTIME_ERROR, OUTPUT_LIMIT_ERROR, COMPILE_ERROR, SYSTEM_ERROR, VALIATOR_ERROR, EXIT_NORMALLY,SYSCALL_RESTRICTION,SEGMENTATION_FAULT};

enum exitStatus programStatus;

int ReadTimeConsumption(pid_t pid);
int ReadMemoryConsumption(pid_t pid);
void updateConsumption(pid_t pid);
void pipeReadToStruct(int file,struct stRunStatus *result);//unused
void pipeWriteFromStruct(int file,struct stRunStatus result);

void timeLimitHandler(int signo);
void memoryLimitHandler(int signo);
void printMemory(char memoryPath[]);

void timeLimitationHandler(int signo)
{
//    printf("Time Limit Error!\n");
//    printf("The pid is %d\n", pid);
    kill(pid, SIGKILL);
}

void memoryLimitHandler(int signo)
{
    printMemory(memoryPath);
}

void printMemory(char memoryPath[])
{
    FILE *f = fopen(memoryPath, "r");
    int i;
    
    if (f == NULL)
    {
//        printf("Open %s Failed!\n", memoryPath);
        return ;
    }
    
    for (i = 0; i < 15; i++)
    {
        fgets(buffer, LENGTH - 1, f);
//        printf("%s", buffer);
    }
    
    fscanf(f, "%s %i", buffer, &i);
//    printf("%s", buffer);
//    printf("\nFinished reading %s\n", memoryPath);
    fclose(f);
}

int main (int argc, char * argv[])
{
    int piped;
    int status;
    struct rlimit executableLimit;

    /*
    float passTime;
    struct timeb before, after;
    struct timeval before1, after1;
    FILE *rf;
    */

    if (argc < 7)
    {
        printf("Argument Error!\n");
        exit(0);
    }
   //addjava
   	int language = atoi(argv[7]);
   //end   
    if(language !=4){
    timeLimitation = atoi(argv[2]);
    memoryLimitation = atoi(argv[3]);
    piped = atoi(argv[8]);  //pipe write port
    }
    else{  
	    /*
	     * comment by: Jialin Wu
	     * don't set limit for java app
    	     * because there exist a problem when set time and memory limitation running a java app
	     */
    timeLimitation = -3;
    memoryLimitation = -4;
    }
  
    sprintf(resultFile, "%s", argv[4]);
    sprintf(inputFile, "%s", argv[5]);
    sprintf(outputFile, "%s", argv[6]);

    if (timeLimitation < 1000)
    {
        timeLimitation = 1;
    }
    else
    {
        timeLimitation = timeLimitation / 1000;
    }

    pid = fork();
    if (pid)
    {
	timeConsumption=0;
	memoryConsumption=0;	
	       
	/*
	 * old code
	  alarm(timeLimitation*10);
          signal(SIGALRM, timeLimitationHandler);
        //sprintf(memoryPath, "/proc/%i/status", pid);
//        ftime(&before);
        //gettimeofday(&before1, NULL);
//        printMemory(memoryPath);
//
	 */
	struct user_regs_struct regs;
	int firstExecute = 1;
	/*
	 * comment by: Jialin Wu
	 * refer to ZOJ,but not completely same with ZOJ.[https://code.google.com/p/zoj/source/browse/trunk/judge_client/client/tracer.c]
	 * judge time exceedance by signal SIGXCPU,if time consumption exceexed,set time consumption to zero and memory consumption to zero.
	 * judge memory exceedance by comparing memory consumption and memoryLimitaion every time getting memory consumption  from proc/$pid/status,if memory consumption is greater than the memoryLimitaion,use ptrace(PTRACE_KILL,pid,NULL,NULL) to stop the user app, and set time consumption to zero and memory consumption to zero.
	 */
	while(waitpid(pid,&status,0) > 0){
			
			/*
			 * old code
	//	while(1)
		//updateConsumption(pid);	
		//wait(&status);
		//programStatus = RUNTIME_ERROR;
		if (WIFEXITED(status))
		{
		   // updateConsumption(pid);	
		    programStatus = EXIT_NORMALLY;
		    break;
	//            printf("The child process %i exit normally!\n", pid);
		}
		else if (WIFSIGNALED(status))
		{
	//            printf("The child process %i exit abnormally\n", pid);
		    if (WTERMSIG(status) == SIGXCPU)
		    {
	//                printf("Time Limit Error !\n");
			programStatus = TIME_LIMIT_ERROR;
			break;
		    }
		    else if (WTERMSIG(status) == SIGSEGV)
		    {
	//                printf("Memory Limit Error !\n");
			programStatus = MEMORY_LIMIT_ERROR;
			break;
		    }
		    else if (WTERMSIG(status) == SIGXFSZ)
		    {
	//                printf("file size ouput  Limit Error !\n");
			programStatus = OUTPUT_LIMIT_ERROR;
			break;
		    }
		}
		*/
		if (WIFSIGNALED(status))
		{
			if(WTERMSIG(status) == SIGKILL){
			programStatus = RUNTIME_ERROR;
			}
			break;
		}
		if(!WIFSTOPPED(status)){
			if (WTERMSIG(status))
			{
			programStatus = RUNTIME_ERROR;
			}	
			break;
		}
		/*
		 * comment by: Jialin Wu
		 *copy from referrence manual: While being traced, the tracee will stop each time a signal is delivered, even if the signal is being ignored. (An exception is SIGKILL, which has its usual effect.) The tracer will be notified at its next call to waitpid(2) (or one of the related "wait" system calls); that call will return a status value containing information that indicates the cause of the stop in the tracee. While the tracee is stopped, the tracer can use various ptrace requests to inspect and modify the tracee. The tracer then causes the tracee to continue, optionally ignoring the delivered signal (or even delivering a different signal instead).
		 */
		int sig =  WSTOPSIG(status);
		if (sig != SIGTRAP)
		{
			updateConsumption(pid);	
			if (sig == SIGXCPU)
			{
			programStatus = TIME_LIMIT_ERROR;
				printf("time exceeded\n");
			//timeConsumption =(timeLimitation)*1000+1;
			timeConsumption =0;
			memoryConsumption = 0;
			}
			else if( sig == SIGXFSZ){
			programStatus = OUTPUT_LIMIT_ERROR;
			}
			/*
			else if( sig == SIGKILL){
			programStatus = RUNTIME_ERROR;
			
			}	
			*/
			else if( sig == SIGILL){
			programStatus = RUNTIME_ERROR;
			}
			else if( sig == SIGSEGV){
			programStatus = SEGMENTATION_FAULT;
			memoryConsumption = 0;
			}
			else{
			programStatus = RUNTIME_ERROR;
			}
			break;
		//	ptrace(PTRACE_SYSCALL, pid, NULL, sig);
		}

		if (ReadMemoryConsumption(pid) >= memoryLimitation)
		{
			programStatus=MEMORY_LIMIT_ERROR;
			//memoryConsumption = memoryLimitation+1;
			ptrace(PTRACE_KILL,pid,NULL,NULL);	
			break;
		}
		ptrace(PTRACE_GETREGS,pid,NULL,&regs);
		/*
		 *comment by: Jialin Wu
		 *regs.orig_rax(in ubuntu_64bit) or regs.orig_eax(int centOS_32bit) will hold the syscall num that was invoked.
		 *
		 *compare with the diabled_syscall array in "disabled_syscall.h" to judge whether the syscall invoked is permitted or not,if the values of disabled_syscall[regs.orig_rax]==1,the syscall is forbidden.
		 */
		/*
		 * for x32
		if (disabled_syscall[regs.orig_eax]==1)
		 */
		if (disabled_syscall[regs.orig_rax]==1)
		{
			if (firstExecute)
			{
				/*
				 *comment by: Jialin Wu
				 * the first execv comes from from parent process
				 * so first exec is permitted to run user app
				 *
				 * consider using GlOG to save  which signal is forbiden
				 */
				firstExecute =0;
			ptrace(PTRACE_SYSCALL,pid,NULL,NULL);	
			}
			else{
				//the app is do somethin evil,so kill it! 
			programStatus =SYSCALL_RESTRICTION;
			ptrace(PTRACE_KILL,pid,NULL,NULL);	
			break;
			}

		}	
		if (regs.orig_rax == SYS_exit ||regs.orig_rax ==SYS_exit_group)
		{
			/*
			 * comment by:Jialin Wu
			 * if the user app is exit normally,it will call SYS_exit or SYS_exit_group
			 * detect these two syscall to judge the app is exit normally or not
			 */

			updateConsumption(pid);
			programStatus = EXIT_NORMALLY;

		}
		ptrace(PTRACE_SYSCALL,pid,NULL,NULL);
		}
		/*
		 * old code
	//        printMemory(memoryPath);

	//        ftime(&after);
		//gettimeofday(&after1, NULL);
		tmp = ((long long)after.time)*1000 + ((long long)after.millitm) -
		      ((long long)before.time)*1000 - ((long long)before.millitm);

		printf("RUNNING Time estimated by ftime function %lldms\n", tmp);
		tmp = ((long long)after1.tv_sec)*1000*1000 + 
		      ((long long)after1.tv_usec) -
		      ((long long)before1.tv_sec)*1000*1000 - 
		      ((long long)before1.tv_usec);

	*/
	//        printf("RUNNING Time estimated by gettimeofdaty function %lldus\n", tmp);
		//getrusage(RUSAGE_CHILDREN, &executableResourceUsage);
	/*        

		printf("Total Memory used by %d is %ld KB \n", pid, 
			executableResourceUsage.ru_maxrss);
		*/
		/*
		 * replace with pipe comunication
		 *
		rf = fopen(resultFile, "w");
		if (rf != NULL)
		{
			 *comment by: Jialin Wu
			 * write time,memory consumption and programStatus to file
		    fprintf(rf, "%lld %ld %d", tmp, 
			    executableResourceUsage.ru_maxrss, (int)programStatus);
		    
		    fprintf(rf, "%ld %ld %d", timeConsumption, 
			    memoryConsumption, (int)programStatus);
		    fclose(rf);
		}
		*/
		struct stRunStatus runStatus;	
		runStatus.timeUsed = timeConsumption;
		runStatus.memoryUsed = memoryConsumption;
		runStatus.runExitStatus = programStatus;
		pipeWriteFromStruct(piped,runStatus);

	}
    else
  	{
		ptrace(PTRACE_TRACEME,0,NULL,NULL);
		freopen(inputFile, "r", stdin);
		freopen(outputFile, "w", stdout);
		if (getrlimit(RLIMIT_FSIZE, &executableLimit) == 0)
		{
	/*
		     printf("Soft Limit for virtual memory : %ld\n",
			       executableLimit.rlim_cur);
		     printf("Hard Limit for virtual memory : %ld\n",
			       executableLimit.rlim_max);
	*/

			/*
			 *comment by: Jialin Wu
			 *output limitation is set the const value: 5MB
			 */
		     executableLimit.rlim_cur = 5 * 1024* 1024;  //MB
		     if (setrlimit(RLIMIT_FSIZE, &executableLimit) == 0)
		     {
		//         printf("Set New Limit for file size output limit  successed!\n");
		     }
		}
	 
		if (getrlimit(RLIMIT_AS, &executableLimit) == 0)
		{
		
		/*
		 *comment by: Jialin Wu
		 *I am not setting RLIMIT_AS, I don't think it's necessary,because now using ptrace to trace the excutation.

		     printf("Soft Limit for virtual memory : %ld\n",
			       executableLimit.rlim_cur);
		     printf("Hard Limit for virtual memory : %ld\n",
			       executableLimit.rlim_max);

		*/
		     executableLimit.rlim_cur = -1;  //kb
		     //executableLimit.rlim_cur = 2 * memoryLimitation * 1024;  //unit:kb
		     if (setrlimit(RLIMIT_AS, &executableLimit) == 0)
		     {
	//                 printf("Set New Limit for virtual memory successed!\n");
		     }
		}
		if (getrlimit(RLIMIT_CPU, &executableLimit) == 0)
		{
		/*
		     printf("Soft Limit for cpu time : %ld\n",
			       executableLimit.rlim_cur);
		     printf("Hard Limit for cpu time : %ld\n",
			       executableLimit.rlim_max);
		*/
		     executableLimit.rlim_cur = timeLimitation*2;  //unit:second
		     if (setrlimit(RLIMIT_CPU, &executableLimit) == 0)
		     {
		//                 printf("Set New Limit for cpu time successed!\n");
		     }
		}
		//addjava
		if (4 == language )
		{
			//-cp means specifying the class path
		execl("/usr/java/bin/java","/usr/java/bin/java","-cp",argv[1],"Main",(char *) NULL);
		}
		else{
		execl(argv[1], " ", (char *) NULL);
		}
		//end
	}
    return 0;
}
int ReadTimeConsumption(pid_t pid){
	char buffer[64];
	sprintf(buffer,"/proc/%d/stat",pid);
	FILE* fp = fopen(buffer,"r");
	if (fp == NULL)
	{
		printf("no stat found in proc\n");
		return -1;
	}
	int stime,utime;
	int d1;
	char buff[24],c1;
//	fgetc(fp);
	if (fscanf(fp,"%d %s %c %*d %*d %*d %*d %*d %*u %*u %*u %*u %*u %d %d",&d1,buff,&c1,&utime,&stime) < 2)
	{
		printf("fail read time\n");
		fclose(fp);
		return -1;
		/* code */
	}
//	printf("%d,%s,%c,%d,%d\n",d1,buff,c1,utime,stime );
	fclose(fp);
		static int clktck = 0;
	 if (clktck == 0)
	 {
		 //Inquire about the number of clock ticks per second
		 //the number of clock ticks per second
		 clktck = sysconf(_SC_CLK_TCK); 
	 }
	 //unit: million second
	 return (int) (((double)(stime + utime + 0.0) / clktck) * 1000);
}

int ReadMemoryConsumption(pid_t pid){
	char buffer[64];
	sprintf(buffer,"/proc/%d/status",pid);
	    FILE* fp = fopen(buffer, "r");
	    if (fp == NULL) {   
		   	return -1;
		    }
	    int vmPeak = 0,
		vmSize = 0,
		vmExe = 0,
		vmLib = 0,
		vmStack = 0;
	    

		while (fgets(buffer, 32, fp)) 
		{ 
		 if (!strncmp(buffer, "VmPeak:", 7)) 
		 { 
	      		 sscanf(buffer + 7, "%d", &vmPeak);

       		} else if (!strncmp(buffer, "VmSize:", 7)) 
		{
		       	sscanf(buffer + 7, "%d", &vmSize);

	        } else if (!strncmp(buffer, "VmExe:", 6)) 
		{ 
			sscanf(buffer + 6, "%d", &vmExe);

		} else if (!strncmp(buffer, "VmLib:", 6)) 
		{           
			   sscanf(buffer + 6, "%d", &vmLib);

		} else if (!strncmp(buffer, "VmStk:", 6)) 
		{            
			sscanf(buffer + 6, "%d", &vmStack);

		}

		}
//		printf("VmPeak:%d VmSize:%d\n",vmPeak,vmSize );
		fclose(fp);
		if(vmPeak){
			vmSize = vmPeak;	
		}
//		printf("vmsize:%d,vmExe:%d,vmLib:%d,vmStack:%d\n",
		     //vmSize,vmExe,vmLib,vmStack );
		//unit:kb
		return vmSize - vmExe -vmLib -vmStack;
    }

void updateConsumption(pid_t pid){
	if (ReadMemoryConsumption(pid) > memoryConsumption)
	{
		memoryConsumption = ReadMemoryConsumption(pid);
	}

	if (ReadTimeConsumption(pid) > timeConsumption)
	{
		timeConsumption = ReadTimeConsumption(pid);
	}	

}
/*=================cString with struct======================*/
void pipeReadToStruct(int file,struct stRunStatus *result){
	FILE* pstream = fdopen(file,"r");
	//if (fscanf(pstream,"%s%d",getArr,&getInt) == EOF)
	if (fscanf(pstream,"%d%d%d",&result->timeUsed,&result->memoryUsed,&result->runExitStatus) == EOF)
	{

	}
	//cout<<"in read function statu:"<<(*result).statu<<" time:"<<result->timeConsumption<<" memory:"<<result->memoryConsumption<<endl;
	//printf("demo getInt:%d getArr:%s\n",getInt,getArr );
}

void pipeWriteFromStruct(int file,struct stRunStatus result){
	char cstring[1024];
	sprintf(cstring,"%d %d %d",result.timeUsed,result.memoryUsed,result.runExitStatus);
	FILE* pWStream = fdopen(file,"w");
	fprintf(pWStream, "%s",cstring);

}
/*=================end cString with struct======================*/


