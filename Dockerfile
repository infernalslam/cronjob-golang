FROM golang:1.12.4-stretch as builder
WORKDIR /app
COPY go.* /
RUN go mod download
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:latest
WORKDIR /app
COPY --from=builder app .
EXPOSE 8080
ENTRYPOINT ["./app"]