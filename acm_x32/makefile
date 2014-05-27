s=main.cpp
d=MainExecution
all:
#	clang $(s) -lmysqlcppconn -std=c++11 -o $(d)
#

#	/usr/local/gcc4.9/bin/g++ $(s) -lmysqlcppconn -std=c++11 -o $(d)

	g++ $(s) -lmysqlcppconn -g -std=c++11 -o $(d)
	./$(d)
.phony:clean
clean:
	rm -f $(d)

