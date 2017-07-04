all: 
	go test

b:
	go test -bench=. -benchmem -blockprofile p.out -coverprofile cover.out -memprofile mem.out
# .phony:clean
clean:
	rm -f cover.out p.out mem.out

10:
	for i in {0..$^}; do 
	go test 
	done
