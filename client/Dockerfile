# FROM golang
FROM golang:1.9.1

# Set timezone to Asia/Tokyo
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

RUN apt-get update \
  && apt-get install -y build-essential git curl wget \
                        zlib1g-dev libxml2-dev libxslt1-dev \
                        openssl less libssl-dev libreadline-dev vim unzip

# Install protobuf
# https://github.com/grpc/grpc-docker-library
RUN mkdir -p /tmp/protoc && \
  curl -L https://github.com/google/protobuf/releases/download/v3.0.0/protoc-3.0.0-linux-x86_64.zip > /tmp/protoc/protoc.zip && \
  cd /tmp/protoc && \
  unzip protoc.zip && \
  cp /tmp/protoc/bin/protoc /usr/local/bin && \
  chmod go+rx /usr/local/bin/protoc && \
  cd /tmp && \
  rm -r /tmp/protoc

# Install glide
RUN curl https://glide.sh/get | sh

# Add the local package files to the container's workspace.
COPY ./ /go/src/
COPY ./ /go/src/client/

WORKDIR /go/src/client/

# Install fresh for live reloading
RUN go get github.com/pilu/fresh

# Install packages via glide
RUN glide install
