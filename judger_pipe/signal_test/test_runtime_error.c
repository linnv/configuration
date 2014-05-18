#include <stdio.h>

int main ()
{
   int a[10];
   int i;
    
   for (i = 0; i < 10000; i++)
   {
       a[i] = 1000;
   }

   printf("Program exit normally\n");
   return 0;
}
