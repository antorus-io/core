# Stage: Build
FROM ghcr.io/antorus-io/common/go-builder:1.0.0 AS build

LABEL maintainer="András Barabás <barabasandras1@gmail.com>"
LABEL version="1.0.0"

ARG APPLICATION_ENV=ANONYMOUS_INSTANCE
ARG APPLICATION_NAME=antorus
ARG SERVICE_NAME=core
ENV APPLICATION_ENV=${APPLICATION_ENV}

WORKDIR /srv/${APPLICATION_NAME}

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags='-s -w' -a -installsuffix cgo -o ${SERVICE_NAME} ./
