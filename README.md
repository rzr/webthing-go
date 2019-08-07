# WEBTHING-GO: #

[![GitHub forks](
https://img.shields.io/github/forks/rzr/webthing-go.svg?style=social&label=Fork&maxAge=2592000
)](
https://GitHub.com/rzr/webthing-go/network/
)
[![license](
https://img.shields.io/badge/license-MPL--2.0-blue.svg
)](LICENSE)
[![GoDoc](
https://godoc.org/github.com/rzr/webthing-go?status.svg
)](
http://godoc.org/github.com/rzr/webthing-go
)


## USAGE: ##

To get started look at ["example" directory](./example):

```sh
go version
#| go version go1.10.4 linux/amd64

go get github.com/rzr/webthing-go

cd ~/go/src/github.com/rzr/webthing-go
go run example/simplest-thing.go
#| Listening: :8888

curl http://localhost:8888/properties
#| {"on":false}

curl -X PUT --data '{"on": true}'  http://localhost:8888/properties/on
#| {"on":true}

curl http://localhost:8888/properties/on
#| {"on":true}

curl http://localhost:8888 | jq -M .
#| { ...
#| "title": "ActuatorExample"
#| }
```


To import published module just use this alias:

```go
import (
	webthing "github.com/rzr/webthing-go"
)
```


## RESOURCES: ##

* <https://iot.mozilla.org/framework/>
* <https://github.com/julienschmidt/httprouter>
