FROM alpine:3.6

RUN apk add --no-cache ca-certificates

ADD bin/linux/amd64/coordinator /coordinator
EXPOSE 8080
ENTRYPOINT ["/coordinator"]
