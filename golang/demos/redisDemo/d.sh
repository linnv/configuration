cids=`docker ps -a | awk '{print $1}'`

function stop() {
for v in $cids
do
	echo 'stop '$v
	`docker stop "$v"`
done

echo "stop all containers "$cids
}

function start() {
	
for v in $cids
do
	echo 'start '$v
	`docker start "$v"`
done

echo "start all containers "$cids
}

function pause() {
for v in $cids
do
	echo 'pause '$v
	`docker pause "$v"`
done

echo "pause all containers "$cids
}

function remove() {
for v in $cids
do
	echo 'rm '$v
	`docker rm "$v"`
done

echo "remove all containers "$cids
}

case $1 in
	"0" )
		start
		;;
	"1" )
		stop
		;;
	"2" )
		pause
		;;
	"3" )
		remove
		;;
	* )
		echo "usage: ./xxx.sh  option
option:{
0: start all containers
1: stop all containers
2: pause all containers
3: remove all containers
}"
		;;
esac
