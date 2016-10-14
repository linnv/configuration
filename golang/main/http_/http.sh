body='{
"name":"jialin"
}'
# body='{
# "A":{
# "name":"jialin"
# }
# }'
# body='[{
# "Id": 585,
# "ExchangeIds": [10040]
# }]'


request_url="http://localhost:8089?c=3939"
# request_url="http://localhost:8089"
# request_url="http://localhost:7101/advertisers/add"

curl -v "$request_url" -d "$body" | jq .

# curl -v http://127.0.0.1:8099/debug\?open\=true
