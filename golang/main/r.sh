#!/usr/bin/env bash
for (( i = 0; i < 5; i++ )); do
	 go run redis.go > out.redis$i 2>&1
	 # go run redis.go 
	 # echo -e "$i\n"
done
