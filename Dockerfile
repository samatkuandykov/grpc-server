FROM golang:latest
RUN mkdir /app
COPY . /app/
WORKDIR /app
RUN apt-get update && apt-get install
RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/proto
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/runtime
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/utilities
EXPOSE 8081