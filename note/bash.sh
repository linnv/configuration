[kill pids](https://stackoverflow.com/questions/3510673/find-and-kill-a-process-in-one-line-using-bash-and-regex)

kill $(ps -aux  | grep 8010 | awk '{print $2}')


### for loop c-c to stop 
[
for (( ; ; ))
do
	sh testSh/bcmmt-ivr.sh
done

]
