# WEBTHING-GO WITH KUBERNETES: #

## SERVICE USING MINIKUBE: ###

As an extra, here is a recipe to run a service using K8s's minikube:

```sh
name="webthing-go"
url="https://raw.githubusercontent.com/rzr/webthing-go/master/extra/tools/kube/$name.yml"
kubectl=kubectl
kubectl version || kubectl="sudo kubectl" # Maybe you'd better check your env

minikube version
#| minikube version: v1.3.0

minikube start || minikube logs --alsologtostderr

$kubectl version
#| Client Version: version.Info{Major:"1", Minor:"15", GitVersion:"v1.15.1", ...
#| Server Version: version.Info{Major:"1", Minor:"15", GitVersion:"v1.15.2", ...

$kubectl apply -f "${url}"
#| deployment.extensions/webthing-go created
#| service/webthing-go created
$kubectl get services
#| webthing-go   NodePort    10.100.249.206   <none>        8888:30080/TCP   4s

minikube service ${name} --url
#| http://192.168.99.102:30080
time minikube service ${name}
#| 🎉  Opening kubernetes service default/webthing-go in default browser...
```


### CLUSTER USING KUBEADM: ###

The process can be automatically move between cluster's nodes:

```sh
project="webthing-go"
kubectl=kubectl
command="go run example/simplest-thing.go"

kubectl get nodes
#| NAME      STATUS                     ROLES    AGE   VERSION
#| slave01   Ready                      <none>   94m   v1.15.2
#| master    Ready                      master   99m   v1.15.2

kubectl describe nodes

kubectl get services $project
#| webthing-go   NodePort    10.104.44.241   <none>        8888:30080/TCP   32s

pod=$($kubectl get all --all-namespaces \
  | grep -o "pod/${project}.*" | cut -d/ -f2 | awk '{ print $1}' \
  || echo failure) && echo pod="$pod"

kubectl describe pods "$pod"

ip=$($kubectl describe pod "$pod" | grep 'IP:' | awk '{ print $2 }') && echo "ip=${ip}"

curl -i http://$ip:8888/ | grep "HTTP/1.1 200 OK"

#| webthing-go-58d7cc6657   1         1         1       34m

# Move process to slave host
kubectl drain $HOSTNAME --ignore-daemonsets  --delete-local-data
kubectl cordon $HOSTNAME
kubectl uncordon $hostname
$kubectl get all --all-namespaces | grep $project
ps auxf | grep -5 "$command" | grep -v grep # on slave
#| go run ...

# Moving back to master host
kubectl drain $hostname --ignore-daemonsets  --delete-local-data
kubectl cordon $hostname
kubectl uncordon $HOSTNAME
$kubectl get all --all-namespaces | grep $project
ps auxf | grep -- "$command" | grep -v grep # on master
#| root      1561  0.2  0.2 3474520 48392 ?       Ssl  16:08   0:26 /usr/bin/containerd
#| root     11758  0.0  0.0 107344  5428 ?        Sl   19:33   0:00  \_ containerd-shim -namespace moby ...
#| root     11779  0.0  0.0   5620  2320 ?        Ss   19:33   0:00  |   \_ /usr/bin/make start
#| root     11798  0.1  0.1 1295752 19404 ?       Sl   19:33   0:00  |       \_ go run example/simplest-thing.go
#| root     11924  0.0  0.0 477272  5696 ?        Sl   19:33   0:00  |           \_ /tmp/go-build000000403/b001/exe/simplest-thing
```

## RESOURCES: ##

* <https://github.com/rzr/webthing-iotjs/wiki/Kube>
