#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>

/*
 * 这里是中文注释
 * 这里是中文注释2
 */
int main ()
{
   pid_t pid_for_current_process;
   pid_t pid_for_parent_process; 
   
   pid_for_current_process = getpid();
   pid_for_parent_process = getppid();
   
   printf("My pid = %d\n", pid_for_current_process);
   printf("Parent's pid = %d\n", pid_for_parent_process);
   
   execl("/usr/bin/vi", "vi", "./test.c", NULL);
   
   printf("If the message is printed, the execl system call failed!\n");

   return 0;
}
