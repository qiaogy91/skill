ARG GO_VERSION=1.23.0
FROM golang:${GO_VERSION} AS build
WORKDIR /src

# 将项目文件复制到容器内 /src
COPY . .

# 编译 Go 程序
RUN CGO_ENABLED=0 go build -o /bin/server .

FROM alpine:latest AS final
RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add \
        ca-certificates \
        tzdata \
        && \
        update-ca-certificates


# 将可执行文件和 etc 目录复制到最终镜像
COPY --from=build /bin/server /bin/
COPY --from=build /src/etc /etc

EXPOSE 8080

ENTRYPOINT [ "/bin/server", "start" ]
