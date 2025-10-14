FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
RUN go build

FROM scratch AS release
WORKDIR /app
COPY --from=builder /app/go-httpd-echo ./
EXPOSE 8080
ENTRYPOINT ["./go-httpd-echo"]
