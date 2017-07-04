#!/usr/bin/env bash
p=`pwd`
echo "$p"
for i in */*.go
do
	echo "$i"
	`sed -i '' 's/newDir/demo/' "$p/$i"`
done
