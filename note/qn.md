
###dd

docker run -it --rm mysql mysql -h172.20.199.1 -uroot -p


set encoding=utf-8


bf7f42a7

sync code between svn and git
### QC
- time unit issue ms or seconds: done
- skip duplicated taskID on pulling: done, but optimize by pull differ time
- ES cluster
- ASR usage: done
- check latest qcadmin version
  - bps use type int
- check status=-1 before text converted
### server
- returnDesk: nodemaster: agent node
- ner server return entity type only
- QA may return L4 which kind is entity property
- cache entityL4 to redis

<!-- - handle KG: -->
<!-- 	- remove entity before request QA; -->
<!-- 	- when request KG; -->
- handle relative questions
- handle 3thrid API for question
- cache2redis.py: get similarQuestions from multi-tables

### outcall
get title -> nodeName;keep,  use reqName

### NewKG
QA(low score)->
	KGintention(implement detail: curpos,template, response whether request KG)->
			KG-Answer(old graphic DB)
### binlog
- SmartBinLog (feature/es-storage

### build ES cluster

<!-- none_match_count reset when CM or HM -->

binlog to ES

<!-- welcome of smartserver -->

<!-- update diff SmartOutcall in aliyun -->

binlog per day

put QC together

#websocket
#performance
#gobwas/ws https://github.com/gobwas/ws

use MQ
---
autoAssict to xq; use type

wx: robot config

### issue

enable 1: 开启机器人？

sed 's/{Version}/abcd/g'

filecut
{

sudo apt install gfortran

apt install libtool


edit makefile
		  -lgfortran\
add ldconfig

download ffprobe
		  https://www.johnvansickle.com/ffmpeg/
}


QC:
ms


{

流转节点：有判断条件，根据条件来流转流程
{"nodeId":"10001","nodeType":"Normal","content":"节点的播报话术","conditions":[{"typeId":"1","keywors":"关键词，竖线分割","labelData:":"标注语料，竖线分割","targetNodeId":"10002"},{"typeId":"2","keywors":"关键词，竖线分割","labelData:":"标注语料，竖线分割","targetNodeId":"100003"},{"typeId":"3","keywors":"关键词，竖线分割","labelData:":"标注语料，竖线分割","targetNodeId":"10004"}]}

{"nodeId":"10004","nodeType":"Normal","nodeName":"未听清楚","content":"比较抱歉了，可能是刚才信号不好，没听清您的回答，您能再讲一遍吗|真不好意思，您能重复一下刚才的话么","conditions":[{"typeId":"1","keywors":"不想说|不能","labelData:":"标注语料，竖线分割","targetNodeId":"10002"}]}

孤儿节点：被打断或者事件触发，进入孤儿节点
{"nodeId":"10002","nodeType":"Normal","nodeName":"质疑机器人","content":"这个不重要的啦，我们给您致电呢，主要是做市场调研答谢出行保障，如有任何疑问可咨询阳光保险客服热线 95510","conditions":[]}
{"nodeId":"10003","nodeType":"Normal","nodeName":"询问公司名称","content":"阳光保险是国内的七大保险集团之一，是中国企业 500 强，服务业 100 强，我们也是通过访问活动让更多的人认识了解我们阳光保险，将来有保险需求的时候可想到我们阳光。","conditions":[]}


打断总节点：多种打断条件，会跳转到不同节点
{"nodeId":"90001","nodeType":"Interrupt","content":"","conditions":[{"typeId":"1","keywors":"你.*是.*（人|机器）","labelData:":"","targetNodeId":"10002"},{"typeId":"2","keywors":"阳光保险|没听说|是什么公司|干什么的","labelData:":"标注语料，竖线分割","targetNodeId":"10003"},{"type":"3","keywors":"给别人|领*（保险）|给家人|媳妇*孩子*老人|赠送","labelData:":"标注语料，竖线分割","targetNodeId":"10004"}]}

事件总节点：根据触发事件，会跳转到不同节点
{"nodeId":"90002","nodeType":"Event","content":"","conditions":[{"typeId":"1","keywors":"","labelData:":"","targetNodeId":"10004"},{"typeId":"2","keywors":"","labelData:":"","targetNodeId":"10005"},{"typeId":"3","keywors":"","labelData:":"","targetNodeId":"10005"}]}

}



{

* 52d735b (HEAD -> master) code
* d019c8e (git-svn) fix  old version QC
* 161a5a9 fix struct
* c513eb0 abcd
* bfd42c1 todo
* b20709b todo
* 00818a4 update config
* 88d4d4e smart qc
* 1d1e782 smart qc * e00273c smart qc }

{

质检后台，SmartQC：对接 ccod，asr（不用语音切割成多个 n>2），　功能质检
	- 47.95.6.82  /data/go/src/SmartQC
	- 目前 ccod, 捷通 asr 已经对接完成，
	- 正在对接阿里云 asr
	- 项目地址：[http://192.168.1.199:3000/qnzs/SmartQC.git]

文件切割，FileCutProxy,asr 要求切割语音时使用：对接 SmartQC,asr
	- 接受 SmartQC 的语音文件，切割后，请求 asr 全部转换成文本后组装，向 SmartQC 返回语音文本
	- 未部署，项目地址：[http://192.168.1.199:3000:3000/qnzs/FileCutProxy.git]
}
redis pass
57897a692e6a9b1751df2dd063eba41e


### test env docker
[
docker run --name outcall-web -v /mnt/web/outcall:/data/outcall-web -p 8401:80 -it  outcall-web:1.0
]

docker run --name outcall -v /data/go:/data/go -p 8092:8092 -p 8094:8094 -p 8095:8095 -p 8093:8088 -it bash  smart_adis:0.1
{

docker run --name outcall -v /data/go:/data/go -p 8019:8019 -p 8088:8088 -p 8091:8091 -p 8107:8017 -it  smartcallout:0.1


}
[

serverPort: 8088
redisMain: '172.17.160.39:6379'
redisBack: '172.17.160.39:6380'
#RedisMain:  127.0.0.1: 6379
#RedisBack:  127.0.0.1: 6380

#AddrNlpServer: http://127.0.0.1:9091/
addrNlpServer: 'http://127.0.0.1:8001/OutCall/recognition'
#AddrNlpServer: http://47.95.6.82: 8080/SmartServer/robot
nlpServerTimeoutSecond: 5
nodeLoopMax: 3
reusePreNodeWhenMatchingDefaultNode: 0
sayHelloAfterSeconds: 1

defaultEnterprise: 'boc'
defaultNodeName: '281'
flowIdStr: '@NUMATTEMPTS'
taskIdStr: CaseID
userNameStr: Cust_Name
userPhoneStr: '@NUMBER'

#next(Round-Robin), random
replyMethod: next
enableDebug: 0
endDirectly: 0
]

[

[service]
ServerPort:9086
#RedisMain:47.95.6.82:6379
#RedisBak:47.95.6.82:6380
RedisMain:172.17.160.38:6379
RedisBak:172.17.160.38:6380

MySqlAddress:101.200.56.250:3306
MySqlUser:sa
MySqlPwd:Root@qnzs123

CorpusTargetDir:Outcall:/data/go/src/SmartOutCall/test_data
CorpusTargetDirNlp:/data/go/src/SmartOutCall/test_data/nlp
]

binlog
[

[service]
QuestionClassifyScore:0.4
#间隔解析日志 default 5min
ParseIntervalSeconds:300

MySqlAddress:172.17.160.38:3306
MySqlUser:sa
MySqlPwd:Root@qnzs123

#smartserver 产生的日志文件所在目录
ParseTargetLogDir:/data/go/src/SmartOutCall/log/
#ParseTargetLogDir:/home/jialin/go/src/SmartOutCall/log/

#本服务停运时保存 工作上下文 目录
StateDir:/data/outcallbinlog/

#ParseIntervalSeconds:10
#MySqlAddress:127.0.0.1:3306
#MySqlUser:abc
#MySqlPwd:abc

#binlog database name
DefaultEnterprise:callout
#
#ParseTargetLogDir:/home/jialin/go/src/SmartBinLog/data/
#StateDir:/data/smartbinlog/
]

        RPair_RedisPass      = "QNzs@.root_1347908642"

	http://47.95.6.82:9086/sync?flowId=100000011_371
	  \"query\": \"first|Sex|女士|@NUMATTEMPTS|385|@eID|100000011\",
	    \"query\": \"first|Sex|女士|@NUMATTEMPTS|580|@eID|100000011\",


57897a692e6a9b1751df2dd063eba41e
p
[
(39.98.89.35)
ivr: 8091 http/tcp

nodemaster: 9086 http/tcp
model-outcall: 8001 http/tcp

]


[
main:39.98.89.35

port:usage
8091: ivr LB
8086: outcall LB
8001: outcallmodel LB
9086: nodemaster LB

8208: smartQC
6379:redis main
6380:redis back
9200: es2.x
]

[
back:39.98.93.32

port:usage
8091: ivr entry
8086: outcall entry
8001: outcallmodel entry
9086: nodemaster entry

3306: mysql
9200: es2.x
]

[
test:47.105.196.123

8086: outcall
8001: outcallmodel
9086: nodemaster
3306: mysql
8208: smartQC
9200: es2.x
8080: web
]
