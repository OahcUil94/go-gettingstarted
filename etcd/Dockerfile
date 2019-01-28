FROM alpine:3.8

ARG ETCD_VER

RUN apk add --update ca-certificates openssl tar && \
    wget https://github.com/etcd-io/etcd/releases/download/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz && \
    tar xzvf etcd-${ETCD_VER}-linux-amd64.tar.gz && \
    mv etcd-${ETCD_VER}-linux-amd64/etcd* /bin/ && \
    apk del --purge tar openssl && \
    rm -rf etcd-${ETCD_VER}-linux-amd64* /var/cache/apk/*

ADD run.sh /bin/run.sh
RUN chmod +x /bin/run.sh

VOLUME /data
EXPOSE 2379 2380 4001 7001

# CMD ["/bin/sh", "-c", "/bin/run.sh"]

ENTRYPOINT ["/bin/run.sh"]