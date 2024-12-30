FROM golang:1.23 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/binary main.go

FROM alpine:3
RUN apk update && apk add --no-cache ca-certificates tzdata && update-ca-certificates
COPY --from=builder /app/binary /bin/app

ENTRYPOINT ["/bin/app"]