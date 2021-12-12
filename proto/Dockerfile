FROM golang:1.17-buster

ARG VERSION=3.19.1

RUN apt update -y && apt install -y curl unzip
RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v${VERSION}/protoc-${VERSION}-linux-x86_64.zip
RUN unzip protoc-${VERSION}-linux-x86_64.zip -d /usr/local
RUN rm protoc-${VERSION}-linux-x86_64.zip

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN export PATH="$PATH:$(go env GOPATH)/bin"

ENTRYPOINT ["protoc"]
