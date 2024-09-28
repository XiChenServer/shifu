# 使用官方 Go 镜像
FROM golang:1.20 AS builder

# 设置工作目录
WORKDIR /app

# 复制源代码
COPY . .

# 编译 Go 应用
RUN go build -o measurement-app .

# 使用小型基础镜像
FROM alpine:latest

# 复制编译后的二进制文件
COPY --from=builder /app/measurement-app /measurement-app

# 设置容器入口点
ENTRYPOINT ["/measurement-app"]
