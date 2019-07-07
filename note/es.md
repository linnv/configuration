### es7
docker run -p 9201:9200 -p 9301:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.1.1
docker run -p 9201:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:6.8.0


curl -X GET "localhost:9201/_cat/indices?v"

curl -X GET "localhost:9201/$1/_search?pretty=true""


curl -X GET "localhost:9201/$1/_search?pretty=true"   -H 'Content-Type: application/json' -d '{
  "sort": [
       { "@timestamp" : {"order" : "desc"}}
  ]
}'

curl -X GET "localhost:9201/bcmm-waterlog-2019.06/_search?pretty=true"   -H 'Content-Type: application/json' -d '{
  "sort": [
       { "@timestamp" : {"order" : "desc"}}
  ]
}'

[
curl -XPUT 'http://192.168.1.125:9201/_template/filebeat' -H 'Content-Type: application/json' -d@/home/jialin/appdemo/filebeat-7.1.1-linux-x86_64/bcmmjson.json

]

### ES write policy
last-write-wins

https://www.elastic.co/guide/cn/elasticsearch/guide/current/index-doc.html

### check struct
https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-mapping.html

check version

curl http://47.95.6.82:9200

bat -method=GET "localhost:9200/twitter"


curl -XGET 'http://127.0.0.1:9222/12312312340/12312312340/_search'
curl -X POST "localhost:9222/_reindex" -H 'Content-Type: application/json' -d'
{
 "source": {
 "index": "12312312335",
 "type": "12312312335"
 },
 "dest": {
  "index": "12312312340",
  "type": "12312312340"
  }
}
'

### backup

path.repo: ["/data/es_backup"]

curl -XPUT 'http://127.0.0.1:9200/_snapshot/my_backup' -H 'Content-Type: application/json' -d '{
  "type": "fs",
  "settings": {
     "location": "/data/es_backup",
     "compress": true
  }
}'


curl -XPUT 'http://127.0.0.1:9200/_snapshot/my_backup/snapshot_1?wait_for_completion=true'


curl -XGET 'http://127.0.0.1:9200/_snapshot/my_backup'

curl -XDELETE 'http://127.0.0.1:9200/_snapshot/my_backup'

curl -XGET 'http://127.0.0.1:9200/_snapshot/my_backup/snapshot_1'

curl -XPUT 'http://127.0.0.1:9200/_snapshot/my_backup/snapshot_1/_restore'

curl -XPUT 'http://127.0.0.1:9200/_snapshot/my_backup/snapshot_1/_restore?wait_for_completion=true'

curl -XDELETE 'http://127.0.0.1:9200/_snapshot/my_backup'

curl -XGET 'http://127.0.0.1:9200/_snapshot/my_backup/snapshot_a'
curl -XDELETE 'http://127.0.0.1:9200/_snapshot/my_backup/snapshot_a'

curl -XPOST 'http://127.0.0.1:9200/qc_index/_close'

curl -XPOST 'http://127.0.0.1:9200/_snapshot/my_backup/snapshot_a/_restore?wait_for_completion=true'

curl -XPUT 'http://127.0.0.1:9200/_snapshot/my_backup/snapshot_a?wait_for_completion=true' -d '{
  "indices": "qc_index",
  "ignore_unavailable": true,
  "include_global_state": false
}'

curl -XPUT 'http://127.0.0.1:9222/_snapshot/my_backup/snapshot_a?wait_for_completion=true' -H 'Content-Type: application/json' -d '{
  "indices": "12312312335",
  "ignore_unavailable": true,
  "include_global_state": false
}'

elasticdump --input=http://localhost:9222/12312312335 --output=http://localhost:9222/12312312340 --type=analyzer

elasticdump --input=http://localhost:9222/12312312335 --output=http://localhost:9222/12312312340 --type=mapping

elasticdump --input=http://localhost:9222/12312312335 --output=http://localhost:9222/12312312340 --type=data

# bin/logstash -f logstash-simple.conf
input {
    beats {
        port => "5044"
    }
}
# The filter part of this file is commented out to indicate that it is
# optional.
filter {
    grok {
     #grok 里边有定义好的现场的模板你可以用，但是更多的是自定义模板，规则是这样的，小括号里边包含所有一个 key 和 value，例子：（?<key>value），比如以下的信息，第一个我定义的 key 是 data，表示方法为：?<key> 前边一个问号，然后用<>把 key 包含在里边去。value 就是纯正则了，这个我就不举例子了。这个有个在线的调试库，可以供大家参考，http://grokdebug.herokuapp.com/
        # match => { "message" => "(?<date>\d{4}/\d{2}/\d{2}\s(?<datetime>%{TIME}))\s-\s(?<status>\w{2})\s-\s(?<respond_time>\d+)\.\d+\w{2}\s-\s%{IP:client}:(?<client-port>\d+)\[\d+\]->%{IP:server}:(?<server-port>\d+).*:(?<databases><\w+>):(?<SQL>.*)"}
        match => { "message" => "(?<index>.*),(?<data>.*)"}
#过滤完成之后把之前的 message 移出掉，节约磁盘空间。
        remove_field => ["message"]
    }

  urldecode{
  field=>["data"]
}
}
output {
    stdout { codec => rubydebug }
}


(?<date>\d{4}/\d{2}/\d{2}\s(?<datetime>%{TIME}))\s\[(?<codefile>.*):(?<fileline>[\d]*)\]\[(?<level>.*)\]\s(?<field1>.*)\|(?<field2>.*)
o
2018/11/21 11:30:46.718554 [struct.go:116][INFO] 80|%7C2018-11-23+22%3A59%3A41.720681+%2B0800+CST+m%3D%2B243.002892213


[

# bin/logstash -f logstash-simple.conf
input {
    beats {
        port => "5044"
    }
}
# The filter part of this file is commented out to indicate that it is
# optional.
filter {
    grok {
    #match => { "message" => "(?<date>\d{4}/\d{2}/\d{2}\s(?<datetime>%{TIME}))\s\[(?<codefile>.*):(?<fileline>[\d]*)\]\[(?<level>.*)\]\s(?<ctx>.*)\|\|(?<content>.*)"}
        match => { "message" => "(?<date>\d{4}/\d{2}/\d{2}\s(?<datetime>%{TIME}))\s\[(?<codefile>.*):(?<fileline>[\d]*)\].*35m\[(?<level>\w*)\]\s(?<ctx>.*)\|\|(?<content>.*)"}

        remove_field => ["message"]
    }
}
output {
    stdout { codec => rubydebug }
}
]

### copy index to another

curl -XGET 'http://127.0.0.1:9222/12312312335/12312312335/_search'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/t3'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/tx30'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/tx3'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/295c24496b0011ae:3111330007:040285'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/2a5c24a093001390:3111330007:040261'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/2a5c24972e00135d:3111330007:040261'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/295c2470f80011ee:3111330007:040261'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/2a5c24701f0011e2:3111330007:040285'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/295c25e4d60015e3:3111330007:040285'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/2a5c247c30001261:3111330007:040285'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/2a5c24793b00123f:3111330007:040261'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/2a5c243cbd001111:3111330007:040261'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/295c25e82a0015fc:3111330007:040285'

curl -XDELETE 'http://127.0.0.1:9222/12312312340/12312312340/2a5c24469e001189:3111330007:040285'
curl -XDELETE 'http://127.0.0.1:9222/12312312340/12312312340/295c274b830016f2:3111330007:040261'
curl -XDELETE 'http://127.0.0.1:9222/12312312340/12312312340/2a5c2475e200121b:3111330007:040285'
curl -XDELETE 'http://127.0.0.1:9222/12312312340/12312312340/295c258e250014a2:3111330007:040285'
curl -XDELETE 'http://127.0.0.1:9222/12312312340/12312312340/2a5c246e9b0011d9:3111330007:040285'
curl -XDELETE 'http://127.0.0.1:9222/12312312340/12312312340/295c25796200141b:3111330007:040285'
curl -XDELETE 'http://127.0.0.1:9222/12312312340/12312312340/2a5c2485470012c4:3111330007:040261'
curl -XDELETE 'http://127.0.0.1:9222/12312312340/12312312340/295c24815100129d:3111330007:040285'
curl -XDELETE 'http://127.0.0.1:9222/12312312340/12312312340/295c259a8e0014f2:3111330007:040261'
curl -XDELETE 'http://127.0.0.1:9222/12312312340/12312312340/295c2486000012d3:3111330007:040285'

curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/2a5c24469e001189:3111330007:040285'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/295c274b830016f2:3111330007:040261'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/2a5c2475e200121b:3111330007:040285'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/295c258e250014a2:3111330007:040285'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/2a5c246e9b0011d9:3111330007:040285'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/295c25796200141b:3111330007:040285'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/2a5c2485470012c4:3111330007:040261'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/295c24815100129d:3111330007:040285'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/295c259a8e0014f2:3111330007:040261'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/295c2486000012d3:3111330007:040285'

curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/295c24471d001194:3111330007:040285'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312335/2a5c259be20014f3:3111330007:040261'


curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312340/'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312340/'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312340/'
curl -XDELETE 'http://127.0.0.1:9222/12312312335/12312312340/'


295c24496b0011ae:3111330007:040285'
2a5c24a093001390:3111330007:040261'
2a5c24972e00135d:3111330007:040261'
295c2470f80011ee:3111330007:040261'
2a5c24701f0011e2:3111330007:040285'
295c25e4d60015e3:3111330007:040285'
2a5c247c30001261:3111330007:040285'
2a5c24793b00123f:3111330007:040261'
2a5c243cbd001111:3111330007:040261'
295c25e82a0015fc:3111330007:040285'

2a5c24469e001189:3111330007:040285'
295c274b830016f2:3111330007:040261'
2a5c2475e200121b:3111330007:040285'
295c258e250014a2:3111330007:040285'
2a5c246e9b0011d9:3111330007:040285'
295c25796200141b:3111330007:040285'
2a5c2485470012c4:3111330007:040261'
295c24815100129d:3111330007:040285'
295c259a8e0014f2:3111330007:040261'
295c2486000012d3:3111330007:040285'

295c24471d001194:3111330007:040285'
2a5c259be20014f3:3111330007:040261'


  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0          "task_id" : "2a5c246e9b0011d9:3111330007:040285",
          "task_id" : "295c25796200141b:3111330007:040285",
          "task_id" : "295c24471d001194:3111330007:040285",
          "task_id" : "2a5c259be20014f3:3111330007:040261",
100           "task_id" : "295c24815100129d:3111330007:040285",
93094           "task_id" : "295c259a8e0014f2:3111330007:040261",
 100 93014            "task_id" : "2a5c2485470012c4:3111330007:040261",
100            "task_id" : "2a5c2475e200121b:3111330007:040285",
  80  11.0M   99          "task_id" : "295c258e250014a2:3111330007:040285",
72 --:--          "task_id" : "295c2486000012d3:3111330007:040285",


curl -X POST "localhost:9222/_reindex" -H 'Content-Type: application/json' -d'
{
 "source": {
 "index": "12312312335",
 "type": "12312312335"
 },
 "dest": {
  "index": "12312312340",
  "type": "12312312340"
  }
}
'
[
curl  "http://127.0.0.1:9222/1000002/1000002/2a5c39bf62002fe7:3111330007:040299?pretty=true"
]
curl  "http://127.0.0.1:9222/1000002/1000002/295c394b8f002f77:3111330007:040457?pretty=true"



[

curl -X POST "http://127.0.0.1:9222/1000002/1000002/2a5c3959b3002fc8:3111330007:040299/_update" -H 'Content-Type: application/json' -d'
{
   "doc" : {
      "ai_status": 0
   }
}
'

]

[

curl  "http://127.0.0.1:9222/1000002/1000002/295c394f13002f92:3111330007:040459?pretty=true"
]


[
curl -X POST "localhost:9000/5201314176/5201314176/_update_by_query" -H 'Content-Type: application/json' -d'
{"script":{"inline":"ctx._source.ai_status=-1"}}
'
]

[
curl -X POST "http://127.0.0.1:9000/2018111601/_update_by_query" -H 'Content-Type: application/json' -d'{"script":{"inline":"ctx._source.ai_status=-1"}}'

]

### batch update by query
```
[
curl -X POST "localhost:9200/5201314176/_update_by_query" -H 'Content-Type: application/json' -d'
{
  "script": {
    "source": "ctx._source.ai_status=-1"
  },
  "query": {
    "range": {
      "ai_operate_time":{
      "gte":1548811366,
      "lte":1548811466
      }
    }
  }
}
'
]

[


curl -X POST "localhost:9200/2018111601/_update_by_query" -H 'Content-Type: application/json' -d'
{
  "script": {
    "source": "ctx._source.ai_status=-1"
  },
  "query": {
    "term": {
    "task_id":"",
    }
  }
}
'

]
{
    "query": {
        "range" : {
            "price" : {
                "gte" : 1000,
                "lte" : 2000
            }
        }
    }
}
[
curl -X POST "localhost:9200/2018111601/_update_by_query" -H 'Content-Type: application/json' -d'
{
  "script": {
    "inline": "ctx._source.ai_status=-1"
  }
}
'

]
```
