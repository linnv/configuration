MBVlinkFile=/Users/Jialin/golang/src/ssp_mbv_web

PClinkFile=/Users/Jialin/golang/src/ssp_web
PCoriginFile=/Users/Jialin/backup-ssp/ssp_web

if [[ -L $MBVlinkFile ]]; then
	echo "remove $MBVlinkFile"
	rm $MBVlinkFile
fi

if [ ! -L $PClinkFile ]; then
	if [[ -e $PCoriginFile ]]; then
		ln -s $PCoriginFile $PClinkFile
		echo "link $PCoriginFile to $PClinkFile"
	else
		echo "$PCoriginFile not exist"
	fi
else
	echo "$PClinkFile already exist"
fi
