#!/bin/sh
process_pid=`ps -axu |grep -v 'grep' |egrep "guetojJudger"|awk '{print $2}'`

if [ -z $process_pid ];  then
	echo "guetojJudger not found~"
else
	echo "found guetojJudger and the pid is: $process_pid"
	kill -9 $process_pid
	echo "guetojJudger $process_pid killed"
fi
