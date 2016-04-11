#LIBS="-I /usr/include/mysql/ -L /usr/lib -lmysqlclient -lz -lm"

target=guetoj_scheduler guetoj_judger

guetoj_all: $(target)

guetoj_scheduler: guetoj_scheduler.o
	gcc -o guetoj_scheduler guetoj_scheduler.o -I /usr/include/mysql -L /usr/lib/mysqladd -lmysqlclient -lm

guetoj_judger: guetoj_judger.o
	gcc -o guetoj_judger guetoj_judger.o

guetoj_scheduler.o: guetoj_scheduler.c
	gcc -c  guetoj_scheduler.c -o guetoj_scheduler.o

guetoj_judger.o: guetoj_judger.c
	gcc -c -o guetoj_judger.o guetoj_judger.c
clean:
	rm guetoj_scheduler guetoj_scheduler.o guetoj_judger guetoj_judger.o
