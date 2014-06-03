#!/bin/sh
process_pid=`ps -axu |grep -v 'grep' |egrep "guetojJudger"|awk '{print $2}'`

if [ -z $process_pid ];  then
	echo "guetojJudger not found~"
else
	echo "found guetojJudger and the pid is: $process_pid"
fi
