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

FROM golang:1.12.9 as webthing-go-builder
MAINTAINER Philippe Coval (p.coval@samsung.com)

ENV DEBIAN_FRONTEND noninteractive
ENV LC_ALL en_US.UTF-8
ENV LANG ${LC_ALL}
    
RUN echo "#log: Configuring locales" \
  && set -x \
  && apt-get update -y \
  && apt-get install -y locales \
  && echo "${LC_ALL} UTF-8" | tee /etc/locale.gen \
  && locale-gen ${LC_ALL} \
  && dpkg-reconfigure locales \
  && sync

ENV project webthing-go
ENV project_dir /go/src/github.com/rzr/${project}/
ADD Makefile ${project_dir}
WORKDIR ${project_dir}
RUN echo "#log: Setup system" \
  && apt-get update -y \
  && apt-get install -y make sudo \
  && apt-get clean \
  && sync

ADD . ${project_dir}
WORKDIR ${project_dir}
RUN echo "#log: ${project}: Preparing sources" \
  && set -x \
  && make all \
  && sync

FROM debian:10
ENV project webthing-go
ENV project_dir /go/src/github.com/rzr/${project}/
COPY --from=webthing-go-builder ${project_dir}/simplest-webthing-go /usr/local/bin/
EXPOSE 8888
ENTRYPOINT [ "/usr/local/bin/simplest-webthing-go" ]
CMD []
