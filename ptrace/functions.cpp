/*
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
*/
int ReadTimeConsumption(pid_t pid){
	char buffer[64];
	sprintf(buffer,"/proc/%d/stat",pid);
	FILE* fp = fopen(buffer,"r");
	if (fp == NULL)
	{
		printf("no stat found in proc\n");
		return -1;
		/* code */
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
		 clktck = sysconf(_SC_CLK_TCK); /* code */
	 }
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
		printf("VmPeak:%d VmSize:%d\n",vmPeak,vmSize );
		fclose(fp);
		if(vmPeak){
			vmSize = vmPeak;	
		}
		printf("vmsize:%d,vmExe:%d,vmLib:%d,vmStack:%d\n",
		     vmSize,vmExe,vmLib,vmStack );
		return vmSize - vmExe -vmLib -vmStack;
    }

