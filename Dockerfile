FROM golang:1.19-alpine as builder

WORKDIR /build

COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o go-yao .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/go-yao ./
COPY --from=builder /build/settings.docker.yml ./settings.docker.yml

EXPOSE 7003
ENTRYPOINT ./go-yao --env=docker
