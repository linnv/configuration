#LIBS="-I /usr/include/mysql/ -L /usr/lib -lmysqlclient -lz -lm"

test_longblob: test_mysql_longblob.o
	gcc -o test_longblob test_mysql_longblob.o -I /usr/include/mysql/ -L /usr/lib64/mysql -lmysqlclient -lm

test_mysql_longblob.o: test_mysql_longblob.c
	gcc -c -o test_mysql_longblob.o test_mysql_longblob.c 

clean:
	rm test_longblob test_mysql_longblob.o 
