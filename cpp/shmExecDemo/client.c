/*
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/mman.h>
#include <sys/types.h>
#include <fcntl.h>
#include <sys/stat.h>
*/
#include <semaphore.h>  //sem_t

#include <sys/types.h>
#include <sys/mman.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
 
char buf[10];
//char *ptr;
struct share{

	int count;
	sem_t mutex;
};
struct share *ptr; 

#define OPEN_FLAG O_RDWR|O_CREAT
#define OPEN_MODE 00777
int main()
{
        int fd;
        fd = shm_open("region",OPEN_FLAG, OPEN_MODE);
        //fd = shm_open("./region", O_CREAT | O_RDWR, S_IRUSR | S_IWUSR);
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
sem_init(&ptr->mutex,1,1); // the cent value is not zore: IPC(mutex must be in share region)  or threads communication
		sem_wait(&ptr->mutex);
		ptr->count++;

		sem_post(&ptr->mutex);

        printf("client: count : %d \n", ptr->count);
        return 0;
}
