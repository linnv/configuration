[
        server_name outcall.qn outcall.qnzs.ai www.outcall.qnzs.ai;

        index index.html index.php;
        root  /var/www/html/Callout;

        rewrite  ^/$  /index.html  last;
        rewrite  ^/(src|images)/(.*)$  /dist/$1/$2  last;
        rewrite ^/(.*).(html|js|css)$ /dist/$1.$2 break;
]



### log
sort -k2 -n yourfile
[
log_format timed_combined '$remote_addr - $remote_user [$time_local] '
    '"$request" $status $body_bytes_sent '
    '"$http_referer" "$http_user_agent" '
    '$request_time $upstream_response_time $pipe';

access_log /var/log/nginx/yourdomain.com.access.log timed_combined;
access_log /var/log/nginx/ac.log  timed_combined;
]


### 使用 nginx 提供 ssl 代理
```
保留以前的 ws 服务提供方式不做任何变更，增加一个 nginx 开启 ssl 代理，配置跟常规的ssl配置有一些细微的变化，那就是header会有一些变化，websocket需要指定header：Upgrade和http version：1.1 ，因此我这里给出配置详情：
server {
    listen       443 ssl;
    server_name  your.domain.com;#你的域名，如果没有域名就去掉
    ssl on;
    #ssl_certificate     127.0.0.1.crt;
    #ssl_certificate_key 127.0.0.1.key;
    ssl_certificate     your.domain.com.pem;#这里可以使用 pem 文件和 crt 文件
    ssl_certificate_key your.domain.com.key;
    ssl_session_timeout 5m;
    ssl_session_cache shared:SSL:50m;
    ssl_protocols SSLv3 SSLv2 TLSv1 TLSv1.1 TLSv1.2;
    ssl_ciphers ALL:!ADH:!EXPORT56:RC4+RSA:+HIGH:+MEDIUM:+LOW:+SSLv2:+EXP;

    location / {
        proxy_pass http://127.0.0.1:19808;# 这里换成你想转发的 ws 服务地址即可
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "Upgrade";
        proxy_set_header X-Real-IP $remote_addr;
        }
}
```
