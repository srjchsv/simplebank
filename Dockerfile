FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./main ./cmd/app/main.go 

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /app/main /usr/bin/main
COPY  .env /usr/bin
WORKDIR /usr/bin
EXPOSE 8080
CMD ["./main"]
