FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
RUN go build

FROM scratch AS release
WORKDIR /app
COPY --from=builder /app/go-httpd-echo ./
ENTRYPOINT ["./go-httpd-echo"]
