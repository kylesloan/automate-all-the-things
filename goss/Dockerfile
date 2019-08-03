FROM alpine:3.10

# keep curl as its needed for goss http checks, for whatever reason the default install doesn't install the newest version? installing 0.3.6
RUN apk add curl && curl -fsSL https://goss.rocks/install | GOSS_VER=v0.3.7 sh
WORKDIR /mnt
COPY Dockerfile /
COPY goss.yaml /mnt/

# run validate for RETRY seconds, retrying every SLEEP seconds, we go for RETRY as GCP can take a while to get ready
CMD ["goss","validate","--format","documentation","--sleep","10s","--retry-timeout","240s"]
