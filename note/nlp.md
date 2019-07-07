## SmartServer
answer:{
"common":"ddd",
"k1":"v1",
"k2":"v2",
}

开启机器人　多选开关后，　除了欢迎语外，其他情况才可以推可能选项
选项必须配合实体识别，e.g.: 时间，数字

意图识别未加入

ask_more_type:

next_ask_elementype:

trace_relative_ask: get finnal answer

dbbase:
	t_knowledge_tree: mutiAnsWerList->[{"value":"xxxquestion1","type":"TYPE_LOCATION"},{"value":"xxxquestion2","type":"TYPE_TIME"}]

Anne_http_server
	http://119.23.32.100:8108/SmartModel/predict?query= 你好

requet knowledge hrbid.py
	/www/htdocs/bnlp

load L4 answer to redis
	ssh gsf.qn
	/home/script
	nohup python3 Cache_Answer_To_Redis_cs.py cs qy11110000 &

corpus-training:
	python QN_Gen_Train_Data.py 20180918 10000000051
	python QN_Train_Model.py  20180918 10000000051
	python3 Anne_Http_Server.py 8010 10000000051 20180918
	nohup python3 Anne_Http_Server.py 8010 10000000051 20180918&

	python QN_Gen_Train_Data.py 20180925 100001111
	python QN_Train_Model.py  20180925 100001111
	python3 Anne_Http_Server.py 8245 100001111 20180925

 /usr/bin/python3 /home/jialin/py/QaServer/Anne_Http_Server.py 8245 100001111 20180925
 /usr/bin/python2 /home/jialin/py/QaServer/Anne_Http_Server.py 8245 100001111 20180925

	python QN_Gen_Train_Data.py 20180925 11110000
	python QN_Train_Model.py  20180925 11110000
	python3 Anne_Http_Server.py 8235 11110000 20180925


[
{
  "errcode": 0,
  "errmsg": "success",
  "rep": "666 不明白您想问以下哪个问题：\n1、办理转换时间、n2、转换是什么 \n3、转换费用如何计算、n4、转换能否提交多次申请 \n\n 请回复数字选择",
  "result": {
    "content": "666 不明白您想问以下哪个问题：\n1、办理转换时间、n2、转换是什么 \n3、转换费用如何计算、n4、转换能否提交多次申请 \n\n 请回复数字选择",
    "contentPos": "",
    "contentPre": "666 不明白您想问以下哪个 问题：\n1、办理转换时间、n2、转换是什么 \n3、转换费用如何计算、n4、转换能否提交多次申请 \n\n 请回复数字选择",
    "contentSelect": [
      "办理转换时间",
      "转换是什么",
      "转换费用如何计算",
      "转换能否提交多 次申请"
    ],
    "needIntoAgent": false
  }
}
]


## unsolve problem
complete match but lower score


[
python app.py -model_dir /tmp/english_L-12_H-768_A-12/ -num_worker=4
python3 app.py -model_dir /home/jialin/py/chinese_L-12_H-768_A-12/ -num_worker=4

docker build -t bert-as-service:1.0 -f ./docker/Dockerfile .

docker run --runtime nvidia -dit -p 5555:5555 -p 5556:5556  -v /home/jialin/py/chinese_L-12_H-768_A-12/:/model -t bert-as-service:1.0 1
docker run  -dit -p 5555:5555 -p 5556:5556  -v /home/jialin/py/chinese_L-12_H-768_A-12/:/model -t bert-as-service:1.0 1

dk8s run --runtime nvidia -dit  -v /home/jialin/py/chinese_L-12_H-768_A-12/:/model -t bert-as-service:1.0 1
]

[
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo apt-key fingerprint 0EBFCD88
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable
sudo apt-get update
sudo apt-get install docker-ce
sudo service docker restart
]

[
cp model.ckpt-0.meta bert_model.ckpt-0.meta
]
