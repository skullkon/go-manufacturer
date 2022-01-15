FROM golang:1.17.2-alpine3.14 as builder

COPY /go/src/go-manufacturer/ .

WORKDIR /go/src/go-manufacturer/

COPY . /go/src/go-manufacturer/

RUN go build -o build/go-manufacturer ./cmd/api

FROM alpine
RUN apk update && apk add --no-cache ca-certificates tzdata && update-ca-certificates
#
COPY --from=builder ./go/src/go-manufacturer/build/go-manufacturer /usr/bin/go-manufacturer

EXPOSE 8080 8080

ENTRYPOINT ["/usr/bin/go-manufacturer"]