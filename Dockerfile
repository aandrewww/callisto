FROM golang:1.20-alpine AS builder
RUN apk update && apk add --no-cache make git
WORKDIR /go/src/github.com/forbole/callisto
COPY . ./
RUN go mod download
RUN make build

FROM alpine:latest
WORKDIR /callisto
COPY --from=builder /go/src/github.com/forbole/callisto/build/callisto /usr/bin/callisto

RUN /usr/bin/callisto init --home /callisto/.callisto
RUN /usr/bin/callisto parse genesis-file --genesis-file-path /callisto/.callisto/genesis.json --home /callisto/.callisto

CMD [ "callisto" ]