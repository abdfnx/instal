FROM alpine:latest

RUN apk update && apk upgrade && apk add --no-cache ca-certificates

COPY instal /usr/bin/instal

ENTRYPOINT ["/usr/bin/instal"]
