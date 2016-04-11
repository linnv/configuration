#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>

int main ()
{
    pid_t pid;
    printf("Test fork system call!\n");
    
    pid = fork();
    
    if (!pid)
    {
        printf("I am the child! pid = %d\n", pid);
        exit(0);
    }
    
    if (pid > 0)
    {
        pid = wait(NULL);
        printf("I am the parent!\n");
        printf("The child with pid = %d is terminated!\n", pid);
    }

    return 0;
}
