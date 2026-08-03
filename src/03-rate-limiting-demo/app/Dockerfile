FROM golang:1.21-alpine AS build
WORKDIR /app
COPY main.go .
RUN go mod init rate-limit-demo && go get github.com/redis/go-redis/v9 && go build -o app main.go

FROM alpine:3.19
COPY --from=build /app/app /app
EXPOSE 8082
CMD ["/app"]
