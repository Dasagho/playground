FROM golang:1.20.1-alpine
EXPOSE 8000
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main ./main.go
CMD ["./main"]