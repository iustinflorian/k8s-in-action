FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go build -o app main.go
ENTRYPOINT ["/app/app"]
