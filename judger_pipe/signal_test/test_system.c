#include <stdio.h>
#include <stdlib.h>

#define LENGTH 200

char buffer[LENGTH];

int main (int argc, char * argv[])
{
    if (argc < 2)
    {
        printf("Invalid arguements!\n");
        exit(1);
    }

    sprintf(buffer, "gcc -o a.out %s 2>comiple_error\n", argv[1]);
    
    printf("%s", buffer);
    system(buffer);
    system("./a.out");
    return 0;
}
