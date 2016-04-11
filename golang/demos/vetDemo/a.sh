#!/bin/bash

# go tool vet a.go
export PRJ=`git config --get remote.origin.url | sed 's/^https:\/\///' | sed 's/\.git$//'`
go vet $PRJ
# go vet a.go
echo $?
echo $PRJ


echo "haha"
