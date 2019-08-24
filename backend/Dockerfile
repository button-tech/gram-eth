FROM golang:latest AS builder

RUN mkdir /build
ADD . /build
WORKDIR /build

RUN go build -o bin/main .

FROM debian:latest
COPY --from=builder /build/bin /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

CMD ["/app/main"]