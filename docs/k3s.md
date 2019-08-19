# WEBTHING-GO WITH K3S #

As an extra, here is a recipe to run a service using K3S:

```sh
name="webthing-go"
image="rzrfreefr/${name}:latest"
kubectl="sudo kubectl"

sudo sync
curl -sfL "https://get.k3s.io" | sh - # v0.8.0
sudo systemctl status k3s.service

$kubectl get nodes # Wait "Ready state"
#| {host}   Ready    master   51s   v1.14.4-k3s.1

$kubectl run --generator="run-pod/v1"  "${name}" --image="${image}"

$kubectl get all --all-namespaces | grep "$name"
#| default       pod/webthing-go                       1/1     Running     0          19s
#| default       pod/webthing-go-688f84fc55-df8n4      1/1     Running     0          40s

pod=$($kubectl get all --all-namespaces \
  | grep -o "pod/${name}-.*" | cut -d/ -f2 | awk '{ print $1}' \
  || echo failure) && echo pod="$pod"
$kubectl describe pod "$pod"  | grep 'Status:             Running'
ip=$($kubectl describe pod "$pod" | grep 'IP:' | awk '{ print $2 }') && echo "ip=${ip}"

curl -I "http://${ip}:8888" | grep 'HTTP/1.1'

sudo cat "/proc/$(pidof simplest-webthing-go | cut -d' ' -f1)/smaps"
#| 00fc2000-00fe3000 rw-p 00000000 00:00 0                                  [heap]
#| Size:                132 kB
#| (...)
#| 7fff1e838000-7fff1e859000 rw-p 00000000 00:00 0                          [stack]
#| Size:                132 kB
#| (...)

# Remove if needed
/usr/local/bin/k3s-uninstall.sh
```


## RESOURCES: ##

* <https://k3s.io/>
* <https://rzr.github.io/webthing-go/>
* <https://github.com/rzr/webthing-go/>
