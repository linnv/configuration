#include <signal.h>
#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>



int main(void)
{    
    for (int sig = 1; sig < NSIG; sig++)
    {
        char *str = strsignal(sig);
        printf("%2d -> SIG%s\n", sig, str);
    }

    return 0;
}
