MBVlinkFile=/Users/Jialin/golang/src/ssp_mbv_web
MBVoriginFile=/Users/Jialin/backup-ssp/ssp_mbv_web

PClinkFile=/Users/Jialin/golang/src/ssp_web

if [[ -L $PClinkFile ]]; then
	echo "remove $PClinkFile"
	rm $PClinkFile
fi

if [ ! -L $MBVlinkFile ]; then
	if [[ -e $MBVoriginFile ]]; then
		ln -s $MBVoriginFile $MBVlinkFile
		echo "link $MBVoriginFile to $MBVlinkFile"
	else
		echo "$MBVoriginFile not exist"
	fi
else
	echo "$MBVlinkFile already exists"

fi
