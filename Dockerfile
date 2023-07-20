FROM golang:1.20-alpine AS build-env

ENV PACKAGES make gcc libc-dev linux-headers bash curl git

WORKDIR /rizon

COPY . .

ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.2.1/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.a

RUN apk add --no-cache $PACKAGES && \
    BUILD_TAGS=muslc LINK_STATICALLY=true make install && \
    rm -rf /var/cache/apk/*

FROM alpine:edge

WORKDIR /rizon

COPY --from=build-env /go/bin/rizond /usr/bin/rizond

CMD ["rizond"]