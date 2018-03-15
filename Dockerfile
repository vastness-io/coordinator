From alpine:3.6

RUN apk update && \
apk add ca-certificates && \
rm -rf /var/cache/apk/*

ADD bin/linux/amd64/coordinator /coordinator
EXPOSE 8080
ENTRYPOINT ["/coordinator"]