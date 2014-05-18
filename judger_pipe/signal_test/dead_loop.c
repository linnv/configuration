#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main ()
{
    int i, j;
    unsigned long k, m;
    char * ptr;
    printf("Starting dead loop\n");
    k = 0xFFFF;
    m = 0;
    while (1)
    {
        for (i = 1; i <= 9; i++)
        {
            for (j = 1; j <= i; j++)
            {
//                printf("%4d", i*j);
            }
//            printf("\n");
        }
        k--;
//        ptr = (char *) malloc (sizeof(char) * 2000);
//        memset(ptr, 1, 2000);
//        m += 2000;
//        printf("K = %ld\n", 0xFFFF - k);
//        printf("allocated memory size = %ldKB\n", m/1024);
    }
    
    printf("Dead loop finished!\n");
    return 0;
}
