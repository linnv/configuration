# COMPILER= go  build
#COMPILER= /usr/local/gcc/bin/g++
# FLASGS= 
# FLASGS= -Wall -O

# FIX= -std=c++11 -g -O2
# FIX= -std=c++11 -lmysqlcppconn -ljsoncpp -g -O2

# app=demo
# GO_FILES=$(wildcard ./src/*.cpp)
# HEAD_SRC=$(wildcard ./src/*.h)
# OBJS=$(CPP_SRC:%.cpp=%.o)


all: 
	go test
	# $(COMPILER) -o $(app) 
# %.o: %.cpp 
# 	$(COMPILER) -c $< -o $@ $(FLASGS) $(FIX) 

b:
	go test -bench=. -benchmem -blockprofile p.out -coverprofile cover.out -memprofile mem.out
# .phony:clean
clean:
	rm -f cover.out p.out mem.out

# run:
# 	./$(app)
