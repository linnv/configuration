s=main.cpp
d=guetJudger
#d=MainExecution
all:
#	clang $(s) -lmysqlcppconn -std=c++11 -o $(d)
#

#	/usr/local/gcc4.9/bin/g++ $(s) -lmysqlcppconn -std=c++11 -o $(d)

	g++ $(s) -lmysqlcppconn -std=c++11 -o $(d)
	./$(d)
.phony:clean
clean:
	rm -f $(d)
	rm ./tmp/*	

