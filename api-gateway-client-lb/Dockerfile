# workspace (GOPATH) configured at /go
FROM golang:1.20 as builder

#
RUN mkdir /app
WORKDIR /app

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go build -o main main.go && \
    mv main /

FROM alpine
COPY --from=builder main .

ENTRYPOINT ["/main"]
