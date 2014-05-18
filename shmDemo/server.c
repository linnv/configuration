#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/mman.h>
#include <sys/types.h>
#include <fcntl.h>
#include <sys/stat.h>
#include <semaphore.h>  //sem_t

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
	int pid;
	int i;

	//ftruncate(fd, 10);

	//	setbuff(stdout,NULL);
	while(1){
		pid =fork();
		if (pid)
		{

			fd = shm_open("./region", O_CREAT | O_RDWR, S_IRUSR | S_IWUSR);
			if (fd<0) {
				printf("error open region\n");
				return 0;
			}
			printf("addresss is:%d\n",fd );
			ftruncate(fd, sizeof(struct share));

			ptr = mmap(NULL,sizeof(struct share), PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);

			if (ptr  == MAP_FAILED) {
				printf("error map\n");
				return 0;
			}

			sem_init(&ptr->mutex,1,1); // the cent value is not zore: IPC(mutex must be in share region)  or threads communication
			sem_wait(&ptr->mutex);
			for ( i = 0; i < 2; ++i)
			{
				ptr->count++;;
				//	ptr->signal = i+10;
				printf("paren: count=%d\n",ptr->count);

			}
			sem_post(&ptr->mutex);

			shm_unlink("./region");

			waitpid(pid,0,0);

		}
		else if (pid ==0)
		{
			fd = shm_open("./region", O_CREAT | O_RDWR, S_IRUSR | S_IWUSR);
			if (fd<0) {
				printf("error open region\n");
				return 0;
			}
			printf("addresss is:%d\n",fd );
			ftruncate(fd, sizeof(struct share));


			ptr = mmap(NULL,sizeof(struct share), PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);
			if (ptr  == MAP_FAILED) {
				printf("error map\n");
				return 0;
			}

			sem_init(&ptr->mutex,1,1); // the cent value is not zore: IPC  or threads communication
			sem_wait(&ptr->mutex);
			for ( i = 0; i < 2; ++i)
			{
				ptr->count++;;
				//	ptr->signal = i+10;
				printf("child: count=%d\n",ptr->count); 
			}
			sem_post(&ptr->mutex);
			shm_unlink("./region");

			exit(0);
		}
		else{
			printf("fork error!\n");	
			return 1;
		}
		sleep(1);
	}
	/*
	   for ( i = 0; i < 100; ++i)
	   {
	   printf("i=%d\n",i );
	   ptr->count = i;
	   ptr->signal = i+10;
	//ptr = i;
	if (i ==99)
	{
	i=0;
	}
	sleep(1);
	}
	*/
	return 0;
}
