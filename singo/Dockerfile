# 构建golang运行环境 使用别名：builder
FROM golang:alpine as builder

ADD . /usr/local/go/src/giligili

WORKDIR /usr/local/go/src/giligili

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server

FROM alpine

ENV MYSQL_DSN=""
ENV REDIS_ADDR=""
ENV REDIS_PW=""
ENV REDIS_DB=""
ENV GIN_MODE="release"
ENV LOG_LEVEL="warning"
ENV TZ Asia/Shanghai

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata

RUN mkdir /www

WORKDIR /www

COPY --from=builder /usr/local/go/src/giligili/api_server /usr/bin/api_server

ADD ./conf /www/conf

RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]

