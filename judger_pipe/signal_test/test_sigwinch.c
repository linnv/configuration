#include <signal.h>
#include <stdio.h>

pid_t pid;

int flag = 1;
void test_sigwinch(int signo)
{
    printf("Catch the signal SIGWINCH!\n");
}

void test_alarm(int signo)
{
    printf("Sending the SIGWINCH signal to myself\n");
    kill(pid, SIGWINCH);
}
int main ()
{
   pid = getpid();
   signal(SIGWINCH, test_sigwinch);
   signal(SIGALRM, test_alarm);
   alarm(10);
   while (flag);
   return 0;
}
