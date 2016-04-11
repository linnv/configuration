#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/mman.h>
#include <sys/types.h>
#include <fcntl.h>
#include <sys/stat.h>
 
char buf[10];
//char *ptr;
struct share{

	int count;
	sem_t mutex;
};
struct share *ptr; 

int main()
{
        int fd;
        fd = shm_open("./region", O_CREAT | O_RDWR, S_IRUSR | S_IWUSR);
        if (fd<0) {
                printf("error open region\n");
                return 0;
        }
        ftruncate(fd, 10);
        ptr = mmap(NULL, sizeof(struct share), PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);
        if (ptr  == MAP_FAILED) {
                printf("error map\n");
                return 0;
        }
        while (ptr->count ){
        printf("count : %d  signal: %d\n", ptr->count,ptr->signal);
	sleep(1);
	}
        return 0;
}
