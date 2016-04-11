executabelName=demo
#x GOPATH=`echo $(GOPATH)`
#x GOPATH=$(shell echo $(GOPATH))
# GOPATH=/home/jialin/golang
# goLibFiles=$(wildcard $(GOPATH)/src/diy_packag/tmp3/*.go)
goLibFiles=$(wildcard $(GOPATH)/src/local.com/httphandler/*.go)
#  golibfiles=$(wildcard $(GOPATH)/src/local.com/*/*.go)
all:
	# go build $(GOPATH)/src/local.com/httphandler/handler.go
	go install 
build:${goLibFiles}
	go build  $(goLibFiles)
run:
	go run demo.go
	# $(executabelName)
clean:
	rm $(GOPATH)/bin/demo
	rm $(GOPATH)/pkg/local.com/*
