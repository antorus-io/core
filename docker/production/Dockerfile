# Stage: Build
FROM ghcr.io/antorus-io/common/go-builder:1.0.0 AS build

LABEL maintainer="András Barabás <barabasandras1@gmail.com>"
LABEL version="1.0.0"

ARG APPLICATION_ENV=ANONYMOUS_INSTANCE
ARG APPLICATION_MODE=DEVELOPMENT
ARG SERVICE=core
ENV APPLICATION_ENV=${APPLICATION_ENV}
ENV APPLICATION_MODE=${APPLICATION_MODE}
ENV SERVICE=${SERVICE}

WORKDIR /srv/antorus

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags='-s' -a -installsuffix cgo -o core ./
