#include <stdio.h>

void say_goodbye(void)
{
   printf("Goodbye from say_goodbye function!\n");
}

int main ()
{
    atexit(say_goodbye);
    printf("I am the main function!\n");
    return 0;
}
