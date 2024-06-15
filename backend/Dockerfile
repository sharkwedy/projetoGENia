FROM golang:1.22.4-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /main main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=build /main .
EXPOSE 8080
CMD ["./main"]
