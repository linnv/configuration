//
//  trace.cpp
//  traceClass
//
//  Created by wujialin on 11/3/14.
//  Copyright (c) 2014 wujialin. All rights reserved.
//

#include "trace.h"
#include "include.h"
#include "functions.h"
#include"atoin.cpp"
int cu = 0;
int BUFFER_SIZE = 100;
int get_memory_usage(int pid, unsigned int *memory_usage_b)
{    int ret;
    FILE *fp = NULL;
    char buffer[BUFFER_SIZE];
    size_t len;
    long size_value;
    snprintf(buffer, BUFFER_SIZE, "/proc/%d/status", pid);
    if ((fp = fopen(buffer, "rb")) == NULL)
    {        ret = -1;
            goto fail;
        }
    while ((fgets(buffer, BUFFER_SIZE, fp)) != NULL)
    {        len = strlen(buffer);
            if ((len >= strlen("VmPeak:\t")) && (strncmp(buffer, "VmPeak:\t", strlen("VmPeak:\t")) == 0))
            { 
	                ret = size_atoin(&size_value, buffer + strlen("VmPeak:\t"), len - strlen("VmPeak:\t") - 1);
	                if (ret != 0)
	                {                ret = -1;
			                goto fail;
			            }
	                if (size_value == 0) 
	                {                ret = -1;
			                goto fail;
			            }
	                *memory_usage_b = size_value;
	                break;
	            }
        }
    ret = 0;
    goto done;
fail:
done:
    if (fp != NULL) fclose(fp);
    return ret;
}

//
void trace::SetLimitInfo(){
    this->trace_ = true;

    limitInfo.tl_ = 1; //seconds
    limitInfo.vml_ = 20*1024;//kbs
    limitInfo.ml_ = 10000;//kbs for code text size;
    limitInfo.sl_ = 1000;//kbs;
    limitInfo.fsl_= 1000;//kbs;
/*
    limitInfo.tl_ = 100; //seconds
    limitInfo.vml_ = 100000;//kbs
    limitInfo.ml_ = 10000;//kbs for code text size;
    limitInfo.sl_ = 1000;//kbs;
    limitInfo.fsl_= 1000;//kbs;
    */
}
int trace::getTime(){
	return this->timeUsed_;
}
int trace::GetMemory(){
	return this->memoryUsed_;
}
bool trace::CreateProcess(){
    int pid = fork();
    if (pid < 0) {
        return false;
    }
    else if( pid > 0){
        this->pid_ = pid;
        return true;
    }
    else{
        //set limit for child
        if (this->limitInfo.tl_ > 0) {
            if (SetrLimit(RLIMIT_CPU, this->limitInfo.tl_) == -1) {
                raise(SIGKILL);
            }
        }
        if (this->limitInfo.vml_ > 0) {
            if (SetrLimit(RLIMIT_AS, this->limitInfo.vml_ * 1024) == -1) {
                raise(SIGKILL);
            }
        }
        
        if (this->limitInfo.ml_ > 0) {
            if (SetrLimit(RLIMIT_DATA, this->limitInfo.ml_ * 1024) == -1) {
                raise(SIGKILL);
            }
        }
        if (this->trace_) {
		printf("info from child: test trace is true \n");
		if(ptrace(PTRACE_TRACEME,0,0,0) == -1){
            //if(ptrace(PTRACE_TRACEME,0,0,0) == -1){
                
                raise(SIGKILL);
            }
        }
        //end setting
        //execl("/bin/ls","ls",NULL);
        execl("./app","app",NULL);
	 //execl("/home/jialin/myTest/traceClass/app","app","",(char *)0); 
	// execl("/usr/java/bin/java","/usr/java/bin/java",
			 //"-Xms32m","-Xmx256m","java",(char *)0);
        printf("executation for app faild\n");
        raise(SIGKILL);
        return false;
    }
}

void trace::updateUsage(){
    int gtime = ReadTimeConsumption(this->pid_);
    int gmemory = ReadMemoryConsumption(this->pid_);
    if (gtime > this->timeUsed_) {
        this->timeUsed_ = gtime;
    }
    if (gmemory > this->memoryUsed_) {
        this->memoryUsed_ = gmemory;
    }
}
void trace::waitForChild(){
    if (!this->CreateProcess() ) {
        printf("create process faild\n");
        raise(SIGKILL);
    }
    printf("child pid:%d\n",this->pid_ );
    int status;int sig;
    unsigned int memory;
    struct rusage ruse;
    for(;;){
	    get_memory_usage(this->pid_,&memory);
	    printf("info from get_memory_usage:%f\n",(memory+0.0)/1024 );
   	this->updateUsage();
       wait4(this->pid_,&status,0,&ruse);
	if( WIFEXITED(status)){
		printf("app exit!\n");	
		break;
			}       
	if(WIFSIGNALED(status)){
	 sig = WTERMSIG(status);
 	switch(sig){
		case 0:
		printf("normally\n");	
		break;	
		case SIGALRM:
			alarm(0);
		case SIGKILL:
		case SIGXCPU:
		printf("time exceed\n");
		break;
		default:
		printf("runtime error!\n");
		break;
	}	 
	break;
	}
           ptrace(PTRACE_CONT,this->pid_,0,0);
           //ptrace(PTRACE_SYSCALL,this->pid_,0,0);
    }
    int time = 
	     (ruse.ru_utime.tv_sec + ruse.ru_stime.tv_sec) * 1000 
	                                 + (double)(ruse.ru_utime.tv_usec + ruse.ru_stime.tv_usec) / 1000;
   // time = (time + 0.0) /1000;
    printf("\ntime used:%d ms\n",time );
    /*
    while (waitpid(this->pid_,&status,0)) {
        if (!WIFSTOPPED(status)) {
         //   this->updateUsage();
            this->result_ = 0;
            printf(" child exited\n");
            break;
        }
        sig = WSTOPSIG(status);
	struct user_regs_struct regs;
	cu++;
        this->updateUsage();
       // ptrace(PTRACE_SYSCALL,this->pid_,0,0);
        if (sig == SIGTRAP) {
		
		printf("int cu:%d\n", cu);
	     ptrace(PTRACE_GETREGS, pid_, 0, &regs);		
            ptrace(PTRACE_SYSCALL,this->pid_,0,0);
            this->updateUsage();

        }
        else{
		printf("int else cu:%d\n", cu);
            ptrace(PTRACE_SYSCALL,this->pid_,0,sig);
	//	ptrace(PTRACE_SYSCALL,pid_,0,0);
	   this->updateUsage();
            //sysCall handle
            //continue;
        }
        
    }
    */
}
