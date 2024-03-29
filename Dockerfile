
FROM golang:1.18-alpine3.16 AS builder
 
WORKDIR /build
RUN adduser -u 10001 -D app-runner
 
ENV GOPROXY https://goproxy.cn
COPY go.mod .
COPY go.sum .
RUN go mod download
 
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o httpserver .
 
FROM alpine:3.16 AS final
 
WORKDIR /app
COPY --from=builder /build/httpserver /app/
#COPY --from=builder /build/config /app/config
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/token.json /app/
COPY --from=builder /build/config.json /app/
COPY --from=builder /build/prompt.txt /app/
 
USER app-runner

ENTRYPOINT ["/app/httpserver"]