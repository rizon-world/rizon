FROM golang:alpine3.13 AS build-env

# Set up dependencies
ENV PACKAGES bash curl make git libc-dev gcc linux-headers eudev-dev python3

WORKDIR /rizon

COPY go.mod .
COPY go.sum .

COPY . .

RUN apk add --no-cache $PACKAGES && make install

FROM alpine:edge

RUN apk add --update ca-certificates

WORKDIR /rizon

COPY --from=build-env /go/bin/rizond /usr/bin/rizond

CMD ["rizond"]
