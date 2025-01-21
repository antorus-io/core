# Stage: Test
FROM ghcr.io/antorus-io/common/go-builder:1.0.0 AS test

LABEL maintainer="András Barabás <barabasandras1@gmail.com>"

CMD ["go", "test", "-cover", "-coverprofile", "./coverage/coverage.out", "-v", "./..."]

# Stage: Coverage
FROM ghcr.io/antorus-io/common/go-builder:1.0.0 AS coverage

LABEL maintainer="András Barabás <barabasandras1@gmail.com>"

WORKDIR /srv/antorus

COPY ./docker/test/entrypoints/docker-entrypoint.sh /

RUN chmod +x /docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]

CMD ["go", "tool", "cover", "-func", "./coverage/coverage.out"]
