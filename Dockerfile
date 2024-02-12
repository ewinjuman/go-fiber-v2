FROM golang:1.22.0-alpine3.19 AS builder
LABEL stage-go-fiber-v2=builder

RUN apk update && apk add --no-cache git
WORKDIR /go/src/go-fiber-v2

COPY . .

#RUN go mod vendor
RUN go build -mod=readonly -o go-fiber-v2
RUN rm -rf vendor

FROM alpine:latest

ENV environment=local

RUN apk add --no-cache tzdata

RUN mkdir /app
WORKDIR /app

RUN apk add busybox-extras

EXPOSE 8888

COPY --from=builder /go/src/go-fiber-v2/go-fiber-v2 /app
COPY --from=builder /go/src/go-fiber-v2/resource /app/resource

RUN mkdir logs

CMD /app/go-fiber-v2 ${environment}