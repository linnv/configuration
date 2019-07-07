kind: ReplicationController
	scale multi-pods to dynamically



spec:
  type: NodePort
  ports:
  - port: 9091 //visit by localip
    nodePort: 30080 //visit by hostname,not ip

### INSTALL
{
https_proxy=http://192.168.1.69:7777
sudo apt-get update && apt-get install -y apt-transport-https
sudo apt install docker.io
sudo systemctl start docker
sudo systemctl enable docker

https_proxy=http://192.168.1.69:7777  curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add
https_proxy=http://192.168.1.69:7777 apt-add-repository "deb http://apt.kubernetes.io/ kubernetes-bionic main"
https_proxy=http://192.168.1.69:7777  apt-get update
https_proxy=http://192.168.1.69:7777 apt-get install -y kubelet kubeadm kubectl kubernetes-cni
swapoff -a
[
docker pull  mirrorgooglecontainers/kube-apiserver:v1.13.0
docker pull  mirrorgooglecontainers/kube-controller-manager:v1.13.0
docker pull  mirrorgooglecontainers/kube-scheduler:v1.13.0
docker pull  mirrorgooglecontainers/kube-proxy:v1.13.0
docker pull  mirrorgooglecontainers/pause:3.1
docker pull  mirrorgooglecontainers/etcd:3.2.24
docker pull  mirrorgooglecontainers/kubernetes-dashboard-amd64:v1.10.0
docker pull  gcrxio/coredns:1.2.6
]
gcrxio
docker pull  mirrorgooglecontainers/coredns:1.2.6

docker tag mirrorgooglecontainers/kube-apiserver:v1.13.0                  k8s.gcr.io/kube-apiserver:v1.13.0
docker tag mirrorgooglecontainers/kube-controller-manager:v1.13.0         k8s.gcr.io/kube-controller-manager:v1.13.0
docker tag mirrorgooglecontainers/kube-scheduler:v1.13.0                  k8s.gcr.io/kube-scheduler:v1.13.0
docker tag mirrorgooglecontainers/kube-proxy:v1.13.0                      k8s.gcr.io/kube-proxy:v1.13.0
docker tag mirrorgooglecontainers/pause:3.1                               k8s.gcr.io/pause:3.1
docker tag mirrorgooglecontainers/etcd:3.2.24                             k8s.gcr.io/etcd:3.2.24
docker tag gcrxio/coredns:1.2.6 k8s.gcr.io/coredns:1.2.6
docker tag mirrorgooglecontainers/kubernetes-dashboard-amd64:v1.10.0 k8s.gcr.io/kubernetes-dashboard-amd64:v1.10.0

kubeadm init --pod-network-cidr=10.17.0.0/16 --service-cidr=10.18.0.0/24
curl -O https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
kubectl apply -f kube-flannel.yml

kubeadm reset
sudo apt-get purge kubeadm kubectl kubelet kubernetes-cni kube*
sudo apt-get autoremove
sudo rm -rf ~/.kube
}


### proxy
[
kubectl proxy --address='0.0.0.0' --accept-hosts='^*$'
]





###microk8s

https://github.com/ubuntu/microk8s/#deploy-behind-a-proxy
[
Your containerd-env file should be only those two lines:

HTTP_PROXY="YOUR_HTTP_PROXY"
HTTPS_PROXY="YOUR_HTTPS_PROXY"
]
sudo systemctl restart snap.microk8s.daemon-containerd.service


get login token

microk8s.kubectl -n kube-system get secret
microk8s.kubectl -n kube-system describe secret default-token-{xxxxx}



### configmap
kubectl create configmap smoc-config --from-file=config.yaml
kubectl describe configmaps smoc-config
kubectl get configmaps smoc-config -o yaml

