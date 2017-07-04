#!/usr/bin/env bash
p=`pwd`
echo "$p"
for i in */*.go
do
	echo "$i"
	`sed -i '' 's/newDir/demo/' "$p/$i"`
done

for i in *; do
	echo "$i"
	if [[ -d "$i" ]]; then
		echo -e "\tis dir"
		continue
	fi
		echo -e "\tis not dir"
# [[ -d "$i" ]] || dir doesnt' exist,then do something e.g. create dir
# [[ -d "$i" ]] || mkdir ~/.vim;
done
