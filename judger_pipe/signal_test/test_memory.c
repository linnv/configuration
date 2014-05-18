#include <unistd.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/time.h>
#include <sys/wait.h>
#include <sys/timeb.h>
#include <sys/resource.h>

#define LENGTH 1024
int timeLimitation;
int memoryLimitation;
struct rusage executableResourceUsage;

pid_t pid, subPid;     
int pi[2];
char memoryPath[LENGTH];
char buffer[LENGTH];
char resultFile[LENGTH];
char inputFile[LENGTH];
char outputFile[LENGTH];

enum exitStatus {COMPILING = 100000, ACCEPTED, PRESENTATION_ERROR, TIME_LIMIT_ERROR, MEMORY_LIMIT_ERROR, WRONG_ANSWER, RUNTIME_ERROR, OUTPUT_LIMIT_ERROR, COMPILE_ERROR, SYSTEM_ERROR, VALIATOR_ERROR, EXIT_NORMALLY};

enum exitStatus programStatus;

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
    int status;
    float passTime;
    struct rlimit executableLimit;

    struct timeb before, after;
    struct timeval before1, after1;
    long long tmp;
    FILE *rf;

    if (argc < 7)
    {
        printf("Argument Error!\n");
        exit(0);
    }
    
    timeLimitation = atoi(argv[2]);
    memoryLimitation = atoi(argv[3]);
    
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
        //pipe(pi);
/*        subPid = fork();
        
        if (subPid == 0)
        {
        }
        else
        {
        }
*/
        sprintf(memoryPath, "/proc/%i/status", pid);
        alarm(timeLimitation*10);
        signal(SIGALRM, timeLimitationHandler);
//        ftime(&before);
        gettimeofday(&before1, NULL);
//        printMemory(memoryPath);
        wait(&status);
        if (WIFEXITED(status))
        {
            programStatus = EXIT_NORMALLY;
//            printf("The child process %i exit normally!\n", pid);
        }
        else if (WIFSIGNALED(status))
        {
//            printf("The child process %i exit abnormally\n", pid);
            if (WTERMSIG(status) == SIGXCPU)
            {
//                printf("Time Limit Error !\n");
                programStatus = TIME_LIMIT_ERROR;
            }
            else if (WTERMSIG(status) == SIGSEGV)
            {
//                printf("Memory Limit Error !\n");
                programStatus = MEMORY_LIMIT_ERROR;
            }
        }
//        printMemory(memoryPath);

//        ftime(&after);
        gettimeofday(&after1, NULL);
/*
        tmp = ((long long)after.time)*1000 + ((long long)after.millitm) -
              ((long long)before.time)*1000 - ((long long)before.millitm);

        printf("RUNNING Time estimated by ftime function %lldms\n", tmp);
*/
        tmp = ((long long)after1.tv_sec)*1000*1000 + 
              ((long long)after1.tv_usec) -
              ((long long)before1.tv_sec)*1000*1000 - 
              ((long long)before1.tv_usec);

//        printf("RUNNING Time estimated by gettimeofdaty function %lldus\n", tmp);
        getrusage(RUSAGE_CHILDREN, &executableResourceUsage);
/*        
        passTime = (executableResourceUsage.ru_utime.tv_sec + 
                   executableResourceUsage.ru_stime.tv_sec) * 1000 +
                   (float)(executableResourceUsage.ru_utime.tv_usec + 
                   executableResourceUsage.ru_stime.tv_usec) / 1000;
*/
/*
        printf("ru_utime.tv_sec = %i\n", executableResourceUsage.ru_utime.tv_sec);
        printf("ru_utime.tv_usec = %i\n", executableResourceUsage.ru_utime.tv_usec);
        printf("ru_stime.tv_sec = %i\n", executableResourceUsage.ru_stime.tv_sec);
        printf("ru_stime.tv_usec = %i\n", executableResourceUsage.ru_stime.tv_usec);
        printf("Total Time used by %d is %f MS\n", pid, passTime);

        printf("Total Memory used by %d is %ld KB \n", pid, 
                executableResourceUsage.ru_maxrss);
*/
        
        rf = fopen(resultFile, "w");
        if (rf != NULL)
        {
            fprintf(rf, "%lld %ld %d", tmp, 
                    executableResourceUsage.ru_maxrss, (int)programStatus);
            fclose(rf);
        }
    }
    else
    {
        freopen(inputFile, "r", stdin);
        freopen(outputFile, "w", stdout);

        if (getrlimit(RLIMIT_AS, &executableLimit) == 0)
        {
/*
             printf("Soft Limit for virtual memory : %ld\n",
                       executableLimit.rlim_cur);
             printf("Hard Limit for virtual memory : %ld\n",
                       executableLimit.rlim_max);
*/
             executableLimit.rlim_cur = memoryLimitation * 1024;
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
             executableLimit.rlim_cur = timeLimitation*2;
             if (setrlimit(RLIMIT_CPU, &executableLimit) == 0)
             {
//                 printf("Set New Limit for cpu time successed!\n");
             }
        }
        execl(argv[1], " ", (char *) NULL);
    }
    return 0;
}
