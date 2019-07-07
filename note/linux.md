lsof -i -P -n | grep LISTEN | grep 8401

```/etc/security/limits.conf
root soft nofile 1655350
root hard nofile 1655350
* soft nofile 955250
* hard nofile 955350
```
these two number must be larger than hard [* soft nofile 955250/* hard nofile 955350[]
```/etc/sysctl.conf
fs.nr_open=2655350
fs.file-max=2655350
```
#### process running time
 ps -o etime= -p 32300


yum list installed



### mount disk
fdisk -l
mount /xx/sba /dir


### show progress bar on copy
rsync --info=progress2 source dest


###find dir only
find . -type d -name "vbox"
