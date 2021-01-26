# stage 1 Generate Tendermint Binary
FROM golang:1.15-alpine as builder
RUN apk update && \
  apk upgrade && \
  apk --no-cache add make

WORKDIR /atlas
COPY . .
# WORKDIR atlas
RUN make build

# stage 2
FROM alpine:3.9
LABEL maintainer="hello@tendermint.com"

RUN apk update && \
  apk upgrade && \
  apk --no-cache add curl bash

COPY --from=builder atlas/build/atlas /usr/bin/atlas
COPY --from=builder atlas/scripts/publish.sh publish.sh

ENTRYPOINT sh publish.sh
