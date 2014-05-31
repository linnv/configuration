#include <time.h>
#include <unistd.h>
#include <stdio.h>
#include <signal.h>

void alarm_handler(int sig)
{
    printf("Alarm rised!\n");
    alarm(5);
    return ;
}

int main ()
{
    alarm(5);
    signal(SIGALRM, alarm_handler);
    while (1);
    return 0;
}
