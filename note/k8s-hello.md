build image by `docker build -t 'hello:1.0'  ....`
deployment
[
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: helloapp
  namespace: kube-public
spec:
  replicas: 2
  template:
    metadata:
      labels:
        app: helloapp
    spec:
      containers:
      - name: helloapp
        image: 'hello:1.0'
        ports:
        - containerPort: 9091

]

service
[
apiVersion: v1
kind: Service
metadata:
  name: helloapp-svc
  namespace: kube-public
  labels:
    app: helloapp-svc
spec:
  type: NodePort
  ports:
  - port: 9091
    nodePort: 30080
  selector:
    app: helloapp
]



kubectl -n kube-system get secret
kubectl -n kube-system describe secret replicaset-controller-token-kzpmc

kubectl -n kube-system describe secret replication-controller-token-hzn6s

kubectl -n kube-system describe secret replicaset-controller-token-xqpgv

kubectl -n kube-system describe secret kubernetes-dashboard-token-r2h8z

kubectl -n kube-system describe secret admin-token-k7z2r


kubectl proxy --address='192.168.1.125' --port=8001 --accept-hosts='^*$'



[

k8s.gcr.io/heapster-amd64:v1.5.2
k8s.gcr.io/heapster-amd64:v1.5.2
cdkbot/addon-resizer-amd64:1.8.1
cdkbot/addon-resizer-amd64:1.8.1
gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.7
gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.7
gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.7
k8s.gcr.io/kubernetes-dashboard-amd64:v1.8.3
k8s.gcr.io/metrics-server-amd64:v0.2.1
cdkbot/addon-resizer-amd64:1.8.1
k8s.gcr.io/heapster-influxdb-amd64:v1.3.3
k8s.gcr.io/heapster-grafana-amd64:v4.4.3
]

[
#!/bin/bash

images=(
k8s.gcr.io/heapster-amd64:v1.5.2
k8s.gcr.io/heapster-amd64:v1.5.2
cdkbot/addon-resizer-amd64:1.8.1
cdkbot/addon-resizer-amd64:1.8.1
gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.7
gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.7
gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.7
k8s.gcr.io/kubernetes-dashboard-amd64:v1.8.3
k8s.gcr.io/metrics-server-amd64:v0.2.1
cdkbot/addon-resizer-amd64:1.8.1
k8s.gcr.io/heapster-influxdb-amd64:v1.3.3
k8s.gcr.io/heapster-grafana-amd64:v4.4.3
)

for imageName in ${images[@]} ;
do
     echo $imageName
     echo ${imageName##*/}
   microk8s.docker pull $imageName
done
]
