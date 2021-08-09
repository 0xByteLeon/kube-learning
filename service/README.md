## Run alpine for test
```shell
kubectl run -i --tty alpine --image=alpine -- sh

apk add --no-cache curl
```

## Access Service
### ClusterIP
```shell
minikube ssh
curl ip:port
```
### NodePort
```shell
minikube service list
minikube service servicename --url
```