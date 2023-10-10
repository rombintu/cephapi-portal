FROM golang:1.21-alpine
WORKDIR /opt/cephapi-portal
COPY . .
RUN go build -mod=vendor -o ./dist/prog cmd/main.go
CMD ["./dist/prog"]