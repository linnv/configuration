int Setrlimit(int resource, unsigned int limit){
	struct rlimit r;
        r.rlim_max = limit + 1;
        r.rlim_cur = limit;
        if (setrlimit(resource,&r) ==-1)
        {   
	                return -1; 
	        }   
        return 0;

}
struct rlimit Getrlimit(int resource){
	struct rlimit r;
	if (getrlimit(resource,&r) ==-1)
        {   
		
		raise(SIGKILL);
	                //return 0; 
	 }   
	printf("source:%d limit_cur:%d,limit_max:%d\n",resource,r.rlim_cur,r.rlim_max );
        return r;

}

struct SettingForProcess{
	int timeLimit_;  //time for app to run
	int memoryLimit_; //memory for code data
	int vmLimit_;  //memory for app to run
	int stackLimit_;
	int outputLimit_;  //output file number for app
	int beTraced_;


};
int CreateProcessAndDoExecute(const SettingForProcess &setting){
	int pid = fork();
	if (pid < 0)
	{
		return -1;
		/* code */
	}
	else if(pid > 0){
		return pid;
	}
	else{
	if(setting.timeLimit_>0){
		if(Setrlimit(RLIMIT_CPU,setting.timeLimit_) == -1){
		printf("setting for app faild\n");	
		raise(SIGKILL);
		}
		Getrlimit(RLIMIT_CPU);
	}	
	if(setting.memoryLimit_ > 0){
		if(Setrlimit(RLIMIT_DATA,setting.memoryLimit_*1024) == -1){
		printf("setting for app faild\n");	
		raise(SIGKILL);
		}
		Getrlimit(RLIMIT_DATA);
	}	
	if(setting.vmLimit_ > 0){
		if(Setrlimit(RLIMIT_AS,setting.vmLimit_*1024) == -1){
		printf("setting for app faild\n");	
		raise(SIGKILL);
		}
		Getrlimit(RLIMIT_AS);
	}	
	if(setting.stackLimit_ > 0){
		if(Setrlimit(RLIMIT_STACK,setting.stackLimit_*1024) == -1){
		printf("setting for app faild\n");	
		raise(SIGKILL);
		}
		Getrlimit(RLIMIT_STACK);
	}	
	if(setting.outputLimit_ > 0){
		if(Setrlimit(RLIMIT_FSIZE,setting.outputLimit_ * 1024) == -1){
		printf("setting for app faild\n");	
		raise(SIGKILL);
		}
		Getrlimit(RLIMIT_FSIZE);
	}	
	if (setting.beTraced_)
	{
		if ( ptrace(PTRACE_TRACEME,0,0,0) == -1)
		{
		printf("set trace faild\n");	
			raise(SIGKILL);
			/* code */
		}
		/* code */
	}
	printf("\nstart to execute\n");
	execl("/home/jialin/myTest/NewTrace/app","app","",(char *)0);
//	execl("/bin/ls","ls",NULL);
	//execl("/bin/date","date",NULL);
	//execl("./app","app",NULL);
	printf("executation for app faild\n");
	raise(SIGKILL);
	}

	return -1;
}
