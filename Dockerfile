FROM golang:latest
ADD . /go/src/github.com/samatkuandykov/grpc-server
RUN apt-get update && apt-get install
RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/proto
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/runtime
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/utilities
RUN go install github.com/samatkuandykov/grpc-server
ENTRYPOINT ["/go/bin/grpc-server"]
EXPOSE 8081