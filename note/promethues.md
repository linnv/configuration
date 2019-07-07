cpu-percent:
[
100 - (avg by (instance) (irate(node_cpu_seconds_total{job="node",mode="idle"}[5m])) * 100)
]

### Redis

redis_exporter -redis.addr 127.0.0.1:6380,127.0.0.1:6381 -redis.password "QNzs@.root_1347908642" -web.listen-address ":9013"



### config sample
https://github.com/prometheus/prometheus/blob/master/config/testdata/conf.good.yml
