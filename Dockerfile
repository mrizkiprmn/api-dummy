FROM golang:1.18-alpine AS builder

WORKDIR $GOPATH/src/api-dummy/

COPY . .

ENV GOSUMDB=off
COPY go.mod .
COPY go.sum .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/dummy

FROM alpine:3.15

COPY --from=builder /go/bin/dummy /go/bin/dummy

ENTRYPOINT ["/go/bin/dummy"]
