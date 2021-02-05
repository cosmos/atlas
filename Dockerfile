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

RUN apk update && \
  apk upgrade && \
  apk --no-cache add curl bash

COPY --from=builder atlas/build/atlas /usr/bin/atlas
COPY ./scripts/publish.sh /publish.sh

RUN chmod +x /publish.sh

ENTRYPOINT ["/publish.sh"]
