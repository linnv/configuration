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

headerAuth="Authorization:xxoobasicauthorization"
# headerAuth="1Authorization:xxoobasicauthorization"
headerBasic="Basic:xxoobasicBasic"
host="9099"
router="/auth"
request_url="http://localhost:$host$router?c=3939"
# request_url="http://localhost:8089"
# request_url="http://localhost:7101/advertisers/add"

# curl -v "$request_url" -d "$body" | jq .
curl -v "$request_url" -d "$body" -H "$headerAuth" -H "$headerBasic"
