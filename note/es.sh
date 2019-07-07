curl -X POST -i -H 'content-type:application/json' -d '
{ "create" : { "_index" : "test", "_type" : "type1", "_id" : "5" } }
{ "field1" : "value5","f2":"v2" }
{ "create" : { "_index" : "test", "_type" : "type1", "_id" : "4" } }
{ "field1" : "value44" }
{ "create" : { "_index" : "test", "_type" : "type1", "_id" : "3" } }
{ "field1" : "value333" }' http://127.0.0.1:9200/_bulk  

curl -X GET -i http://127.0.0.1:9200/test/type1/_search

curl -X GET -i http://127.0.0.1:9200/qc_index/qc_search/_search

[
curl -X GET "http://47.95.6.82:9200/qc_index/5201314175/_search" -H 'Content-Type: application/json' -d'
{
    "query": {
        "regexp":{
            "task_id": "46.*"
        }
    }
}
'
]
[
curl -X GET "http://localhost:9222/12312312340/12312312340/_search?pretty=true" -H 'Content-Type: application/json' -d'
{
    "query": {
        "term":{
            "agentID": "0"
        }
    }
}
'

]

[
curl -X GET "http://localhost:9222/12312312335/12312312335/_search?pretty=true" -H 'Content-Type: application/json' -d'
{
    "query": {
        "term":{
            "agentID": "0"
        }
    }
}
'
]

[
curl -X GET "http://127.0.0.1:9200/5201314176/5201314176/4061229159895364706?pretty=true"
]




###es6
[
curl -X GET "localhost:9201/bcmm-waterlog-2019.06/_search?pretty=true" -H 'Content-Type: application/json' -d'
{
    "sort" : [
        { "ReqServer.ReqTime" : {"order" : "desc"}}
    ]
}
'

]
