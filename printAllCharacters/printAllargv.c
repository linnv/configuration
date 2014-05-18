#include <stdio.h>
#include <stdlib.h>
int main(int argc, const char *argv[])
{
    for(int i= 0;i< 27;i++)    
    {
        printf("argv[%d] : %s\n",i,argv[i]);
        
    }
    return 0;
}
