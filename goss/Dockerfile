FROM alpine:3.10

# keep curl as its needed for goss http checks
RUN apk add curl && curl -fsSL https://goss.rocks/install | sh
WORKDIR /mnt
COPY Dockerfile /
COPY goss.yaml /mnt/

# run validate for RETRY seconds, retrying every SLEEP seconds, we go for RETRY as GCP can take a while to get ready
CMD ["goss","validate","--format","documentation","--sleep","10","--retry-timeout","180"]
