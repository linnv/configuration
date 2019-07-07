curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 && chmod +x minikube && sudo cp minikube /usr/local/bin/ && rm minikube

HTTP_PROXY="http://192.168.1.12:7777"   HTTPS_PROXY="http://192.168.1.12:7777" minikube start  --docker-env HTTP_PROXY="http://192.168.1.12:7777"  --docker-env HTTPS_PROXY="http://192.168.1.12:7777"


