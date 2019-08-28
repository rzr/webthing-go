#!/bin/echo docker build . -f
# -*- coding: utf-8 -*-
# SPDX-License-Identifier: MIT
#{
# Copyright 2019-present Samsung Electronics France SAS, and other contributors
#
# This Source Code Form is subject to the terms of the MIT Licence
# If a copy of the MIT was not distributed with this file
# You can obtain one at:
# https://spdx.org/licenses/MIT.html
#}

FROM golang:1.12.9-alpine3.10 AS webthing-go-builder
LABEL maintainer "Philippe Coval (p.coval@samsung.com)"

ENV project webthing-go   
ENV project_dir /usr/local/src/github.com/rzr/${project}/

RUN echo "#log: Setup system" \
  && apk update \
  && apk add make sudo \
  && rm -rf /var/cache/apk/* \
  && sync

COPY Makefile ${project_dir}
WORKDIR ${project_dir}
RUN echo "#log: ${project}: Preparing sources" \
  && set -x \
  && make rule/setup/alpine \
  && sync

COPY . ${project_dir}
RUN echo "#log: ${project}: Buidling sources" \
  && set -x \
  && make all \
  && sync

FROM alpine:3.10
ENV project webthing-go
ENV project_dir /usr/local/src/github.com/rzr/${project}/
COPY --from=webthing-go-builder ${project_dir}/simplest-webthing-go /usr/local/bin/
EXPOSE 8888
ENTRYPOINT [ "/usr/local/bin/simplest-webthing-go" ]
CMD []
