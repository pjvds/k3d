FROM haproxy:2.1-alpine

RUN apk -U --no-cache add curl ca-certificates\
    && mkdir -p /etc/confd \
    && curl -sLf https://github.com/kelseyhightower/confd/releases/download/v0.16.0/confd-0.16.0-linux-amd64 > /usr/bin/confd \
    && chmod +x /usr/bin/confd \
    && apk del curl \\
    && mkdir -p /etc/rsyslog.d/ \
    && touch /var/log/haproxy.log \
    ln -sf /dev/stdout /var/log/haproxy.log

COPY templates /etc/confd/templates/
COPY conf.d /etc/confd/conf.d/
COPY start-haproxy /usr/bin/

ENTRYPOINT start-haproxy