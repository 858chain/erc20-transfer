FROM golang:1.12.1-alpine3.9 as builder

WORKDIR /go/src/app

RUN apk add git && apk add make && apk add gcc && apk add libc-dev  \
  && apk add --update gcc musl-dev

ENV GOPROXY=https://goproxy.io
ADD . .

RUN make



FROM alpine:latest

COPY --from=builder /go/src/app/bin/erc20-transfer /

EXPOSE 8001
WORKDIR /

CMD ["/erc20-transfer", "run"]

