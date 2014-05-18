struct UsedInfo{
	int timeUsed;
	int memoryUsed;
};
void UpdateUsedInfo(int pid,UsedInfo & info){
//	printf("update pid:%d\n",pid);
	int t = ReadTimeConsumption(pid);
	int m = ReadMemoryConsumption(pid);
	if (t > info.timeUsed)
	{
		info.timeUsed = t;
		/* code */
	}
	if (m > info.memoryUsed)
	{
		info.memoryUsed = m;
		/* code */
	}
//	printf("in update Used time:%d, Memory Used:%d\n",info.timeUsed,info.memoryUsed );
//	printf("update UsedInfo done!\n");

}
void RunProgram(int pid_){
	struct UsedInfo appUsedInfo;
	appUsedInfo.timeUsed = 0;
	appUsedInfo.memoryUsed = 0;
	int status;
	if (pid_ == -1)
	{
		printf("run app faild\n");
		raise(SIGKILL);
	}
	while(waitpid(pid_,&status,0) > 0){
//		printf("waiting pid:%d\n", pid_);
		
		if(!WIFSTOPPED(status)){
			UpdateUsedInfo(pid_,appUsedInfo);
			printf("app exited\n");
		break;	
		}	
		int sig = WSTOPSIG(status);
		//sigInfo
//		printf("signal info: %s\n",strsignal(sig));
	if (sig ==SIGTRAP)
	{
		UpdateUsedInfo(pid_,appUsedInfo);
//		printf("touch execute app start to run\n");
		ptrace(PTRACE_SYSCALL,pid_,0,0);
	//	ptrace(PTRACE_SYSCALL,pid_,0,sig);
//		ptrace(PTRACE_SYSCALL,pid_,0,sig);
//		ptrace(PTRACE_SYSCALL,pid_,0,sig);
//		continue;
		/* code */
	}
	else{
//		printf("sig is not SIGTRAP\n");

			UpdateUsedInfo(pid_,appUsedInfo);
//		ptrace(PTRACE_CONT,pid_,0,0);
		ptrace(PTRACE_SYSCALL,pid_,0,sig);
		ptrace(PTRACE_SYSCALL,pid_,0,0);
	//	printf("app end running\n");
		continue;
	}
	}
	printf("out of wait\n");
	UpdateUsedInfo(pid_,appUsedInfo);
	printf("timeUsed:%dms,memoryUsed:%dkb\n",appUsedInfo.timeUsed,appUsedInfo.memoryUsed );
}
