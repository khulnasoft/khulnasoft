# Copyright 2021 - KhulnaSoft Authors <admin@khulnasoft.com>
# SPDX-License-Identifier: Apache-2.0

FROM alpine:3.15
LABEL maintainer="khulnasoft <admin@khulnasoft.com>"

ARG TARGETARCH
ARG TARGETVARIANT

RUN addgroup -g 10001 -S khulnasoft \
	&& adduser -u 10000 -S -G khulnasoft -h /home/khulnasoft khulnasoft
RUN apk add -U --no-cache ca-certificates libcap tini bind-tools

COPY ./bin/khulnasoft-linux-$TARGETARCH${TARGETVARIANT:+-$TARGETVARIANT} /opt/khulnasoft/khulnasoft-linux-$TARGETARCH${TARGETVARIANT:+-$TARGETVARIANT}
RUN setcap CAP_NET_BIND_SERVICE=+eip /opt/khulnasoft/khulnasoft-linux-$TARGETARCH${TARGETVARIANT:+-$TARGETVARIANT}
RUN ln -s /opt/khulnasoft/khulnasoft-linux-$TARGETARCH${TARGETVARIANT:+-$TARGETVARIANT} /sbin/khulnasoft

RUN mkdir -p /var/opt/khulnasoft \
	&& mkdir -p /var/www/.cache \
	&& mkdir -p /etc/khulnasoft \
	&& chown khulnasoft: /var/opt/khulnasoft /var/www/.cache /etc/khulnasoft

ENV KHULNASOFT_SERVER_PORT 80
EXPOSE 80 443

HEALTHCHECK --interval=1m --timeout=5s \
  CMD wget -qO- http://localhost:80/healthz || exit 1

ENTRYPOINT ["/sbin/tini", "--", "khulnasoft"]

USER khulnasoft
WORKDIR /home/khulnasoft
