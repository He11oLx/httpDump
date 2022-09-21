FROM --platform=$BUILDPLATFORM golang:1.17-alpine3.15 AS builder
WORKDIR /go/src
COPY ./ ./
RUN go build .

FROM --platform=$BUILDPLATFORM alpine:3.15
COPY --from=builder /go/src/httpDump /usr/bin/httpDump
EXPOSE 80
ENTRYPOINT ["/usr/bin/httpDump"]
