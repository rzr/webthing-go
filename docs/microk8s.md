# WEBTHING-GO WITH MICROK8S #

As an extra, here is a recipe to run a service using Microk8s:

### SERVICE: ###

```sh
name="webthing-go"
url="https://raw.githubusercontent.com/rzr/${name}/master/extra/tools/kube/$name.yml"
export PATH="/snap/microk8s/current/:/snap/bin/:$PATH"
kubectl="microk8s.kubectl"
port=8888
#
sudo sync
sudo snap install microk8s --classic # v1.15.2

microk8s.enable dns
$kubectl apply -f "${url}"
$kubectl describe services/$name 
$kubectl get all --all-namespaces | grep "$name" | grep 'Running'
# Through kube-proxy
port=$($kubectl get svc ${name} -o=jsonpath="{.spec.ports[?(@.port==$port)].nodePort}")
curl -i http://127.0.0.1:${port}/ | grep 'HTTP/1.1'
#| HTTP/1.1 200 OK

# Remove if needed
microk8s.reset
$kubectl delete --all pod
$kubectl delete --all node
$kubectl delete --all services
sudo snap remove microk8s
```


### DEPLOYMENT: ###

```sh
name="webthing-go"
image="rzrfreefr/${name}:latest"
export PATH=/snap/microk8s/current/:$PATH
kubectl="kubectl"
port=8888

sudo sync
sudo snap install microk8s --classic # v1.15.2

sudo systemctl status | grep kube
#| (...)
#| /snap/microk8s/743/kube-apiserver
#| (...)

$kubectl cluster-info
#| Kubernetes master is running at http://localhost:8080

$kubectl get services
#| NAME         TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
#| kubernetes   ClusterIP   10.152.183.1   <none>        443/TCP   15m

$kubectl get nodes # Wait "Ready state"
#| {host}   Ready    master   51s   {version}

microk8s.enable dns

$kubectl run --generator=run-pod/v1 "${name}" --image="${image}"

$kubectl get all --all-namespaces | grep "$name" # ContainerCreating, Running
#| default       pod/webthing-go                       1/1     Running     0          19s

pod=$($kubectl get all --all-namespaces \
  | grep -o "pod/${name}-.*" | cut -d/ -f2 | awk '{ print $1}' \
  || echo failure) && echo pod="$pod"
$kubectl describe pod "$pod" | grep 'Status: * Running'
ip=$($kubectl describe pod "$pod" | grep 'IP:' | awk '{ print $2 }') && echo "ip=${ip}"

curl -qi "http://${ip}:${port}" | grep 'HTTP/1.1'
#| HTTP/1.1 200 OK

sudo cat "/proc/$(pidof simplest-webthing-go | cut -d' ' -f1)/smaps" | grep -2 ']'
#| 00fc2000-00fe3000 rw-p 00000000 00:00 0                                  [heap]
#| Size:                132 kB
#| (...)
#| 7fff1e838000-7fff1e859000 rw-p 00000000 00:00 0                          [stack]
#| Size:                132 kB
#| (...)
```

## RESOURCES: ##

* <https://microk8s.io/>
* <https://rzr.github.io/webthing-go/>
* <https://github.com/rzr/webthing-go/>
