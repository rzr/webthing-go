# WEBTHING-GO WITH KUBERNETES: #

## USAGE: ###

As an extra, here is a recipe to run a service using K8s:

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
