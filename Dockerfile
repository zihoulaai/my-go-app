# 构建阶段
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache sqlite-dev gcc musl-dev
RUN CGO_ENABLED=1 GOOS=linux go build -o myapp .

# 运行阶段（使用轻量级镜像）
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/myapp .
EXPOSE 8080
CMD ["./myapp"]