FROM golang:1.22.3 AS builder

ENV GO111MODULE=on
WORKDIR /build
WORKDIR $GOPATH/src/packages/goginapp/
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/main .

FROM alpine

WORKDIR /root
COPY --from=builder /go/main /root/main

ENV GIN_MODE=release

WORKDIR /root