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
[![Build Status](
https://travis-ci.org/rzr/webthing-go.svg?branch=master
)](
https://travis-ci.org/rzr/webthing-go
)
[![pulls](
https://img.shields.io/docker/pulls/rzrfreefr/webthing-go.svg
)](
https://cloud.docker.com/repository/docker/rzrfreefr/webthing-go
)
[![Automated Builds](
https://img.shields.io/docker/automated/rzrfreefr/webthing-go.svg
)](
https://cloud.docker.com/repository/docker/rzrfreefr/webthing-go/timeline
)
[![Build Status](
https://img.shields.io/docker/build/rzrfreefr/webthing-go.svg
)](
https://cloud.docker.com/repository/docker/rzrfreefr/webthing-go/builds
)
[![Go Report Card](
https://goreportcard.com/badge/github.com/rzr/webthing-go
)](
https://goreportcard.com/report/github.com/rzr/webthing-go
)
[![codebeat badge](
https://codebeat.co/badges/f6061081-0f1b-4791-9bef-b439eb379cbc
)](
https://codebeat.co/projects/github-com-rzr-webthing-go-master
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
