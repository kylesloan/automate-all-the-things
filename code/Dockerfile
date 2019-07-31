FROM alpine:3.10

# this file needs to be generated via golang, go build -o code.bin main.go
COPY code.bin /
ENTRYPOINT ["/code.bin"]
EXPOSE 8081

# in the event someone wants to know how this container was built and doesn't have the original Dockerfile
COPY Dockerfile /
