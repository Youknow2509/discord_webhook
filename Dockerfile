# Build go project
FROM golang:alpine AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o crm.dwebhook.com ./cmd/server

# Get certificate
FROM alpine:latest AS cert
RUN apk --no-cache add ca-certificates
RUN update-ca-certificates

# optimize the image
FROM scratch
COPY ./environment /environment
COPY --from=builder /build/crm.dwebhook.com /
COPY --from=cert /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080
ENTRYPOINT [ "/crm.dwebhook.com"]