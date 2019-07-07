docker build -t mytomcat . --no-cache

dk8s ps -a| grep Exited | awk '{print $1}' |xargs microk8s.docker rm




 dk8s rm smartoutcall && dk8s run --name smartoutcall -v /home/jialin/py/SmartOutCallModel:/data/py/SmartOutCallModel -it nlp-py:1.0

curl http://qn:30081/OutCall/testnode?nodeId=1&flowId=281
curl http://192.168.1.69:30081/OutCall/testnode?nodeId=1&flowId=281
curl http://192.168.1.69:30081/OutCall/recognition?nodeId=1&query= 不是 &flowId=281

wrk -t50 -c200  -d30s http://192.168.1.69:30081/OutCall/recognition?nodeId=1&query=%E4%B8%8D%E6%98%AF&flowId=281

wrk -t50 -c200  -d30s http://192.168.1.69:30081/OutCall/recognition%3FnodeId%3D1%26query%3D%E4%B8%8D%E6%98%AF%26flowId%3D281%0D%0A
wrk -t50 -c200  -d30s http://10.152.183.60:30081/OutCall/recognition%3FnodeId%3D1%26query%3D%E4%B8%8D%E6%98%AF%26flowId%3D281%0D%0A

curl http://smartoutcall-svc.kube-public:30081/OutCall/testnode?nodeId=1&flowId=281

wrk  --latency -t100 -c2000  -d20s http://192.168.1.69:30081/OutCall/recognition\?nodeId\=1\&query\=%E4%B8%8D%E6%98%AF\&flowId\=281


ExecStart=/usr/bin/dockerd -H fd:// -g /mnt/docker



yum list installed


###expose
EXPOSE 指令是声明运行时容器提供服务端口，这只是一个声明，在运行时并不会因为这个声明应用就会开启这个端口的服务。在 Dockerfile 中写入这样的声明有两个好处，一个是帮助镜像使用者理解这个镜像服务的守护端口，以方便配置映射；另一个用处则是在运行时使用随机端口映射时，也就是 docker run -P 时，会自动随机映射 EXPOSE 的端口。

要将 EXPOSE 和在运行时使用 -p 《宿主端口》:<容器端口> 区分开来。-p，是映射宿主端口和容器端口，换句话说，就是将容器的对应端口服务公开给外界访问，而 EXPOSE 仅仅是声明容器打算使用什么端口而已，并不会自动在宿主进行端口映射。
