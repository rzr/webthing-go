#! /usr/bin/make -f
# -*- makefile -*-
# ex: set tabstop=4 noexpandtab:
# Copyright: 2019-present Samsung Electronics Co., Ltd. and other contributors
# SPDX-License-Identifier: MPL-2.0

default: all
	@echo "# log: $@: $^"

port?=8888
url?=http://localhost:${port}
main_src?=example/simplest-thing.go
sudo?=sudo

all: get version build
	@echo "# log: $@: $^"

get:
	go $@ || echo "warning: Error is ignored (TODO)"

build:
	go $@

run: ${main_src}
	-go get
	go run $<

start: ${main_src}
	go run $<

client/put:
	@echo
	curl -X PUT -d '{ "on": true}' -i ${url}/properties/on
	@echo

client/get:
	@echo
	curl -i ${url}/properties/on
	@echo

client/get/json:
	@echo
	curl "${url}" | jq -M . \
|| curl -i "${url}"
	@echo

client:
	curl -i ${url}
	@echo

	@echo
	curl -i ${url}/
	@echo

	@echo
	curl -i ${url}/properties
	@echo

	@echo
	curl -i ${url}/properties/on
	@echo

	@echo
	curl -X PUT -d '{ "on": true}' -i ${url}/properties/on
	@echo

	@echo
	curl -i ${url}/properties/on
	@echo

	@echo
	curl -X PUT -d '{ "on": false}' -i ${url}/properties/on
	@echo

	@echo
	curl -i ${url}/properties/on
	@echo

	@echo
	curl -i ${url}/properties/level
	@echo


rule/setup/debian:
	${sudo} apt-get update
	${sudo} apt-get install -y \
 apt-transport-https \
 make \
 curl \
 git \
 golang \
 #EOL

version:
	${MAKE} --version
	go version
