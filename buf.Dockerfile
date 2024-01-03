FROM golang:1.21-alpine3.19 as golang

RUN go install github.com/golang/protobuf/protoc-gen-go@v1.5.2

WORKDIR /workspace

COPY api ./api

FROM bufbuild/buf:1.28.1 as buf

COPY --from=golang /go/bin/protoc-gen-go /usr/local/bin/

ENTRYPOINT ["buf"]

