FROM golang:1.11.2 as builder

WORKDIR /root/go/src/github.com/mirzakhany/gorender
ADD . /root/go/src/github.com/mirzakhany/gorender/
RUN make build

FROM alpine
COPY --from=builder /root/go/src/github.com/mirzakhany/gorender/dist/gorender .
RUN apk add --no-cache ca-certificates && rm -rf /var/cache/apk/*
ENTRYPOINT ["/gorender"]