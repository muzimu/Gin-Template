# 基础镜像
FROM golang:1.22.4 as builder

LABEL maintainer="muzimu"

WORKDIR /app

COPY ../go.mod go.sum ./
# 设置GOPROXY环境变量来更改Go下载源
ENV GOPROXY=https://goproxy.cn
RUN go mod download && go mod verify

COPY ../. .
RUN go build -o app main.go

FROM registry.cn-shanghai.aliyuncs.com/muzimu/alpine:latest

RUN apk add gcompat

EXPOSE 80

WORKDIR /root/
COPY --from=builder /app/app .
RUN chmod +x /root/app

ENTRYPOINT [ "/root/app" ]