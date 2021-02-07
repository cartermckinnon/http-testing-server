FROM golang:alpine AS builder
WORKDIR /app
ADD go.mod go.mod
ADD go.sum go.sum
ADD main.go main.go
RUN go build

FROM alpine
COPY --from=builder /app/http-testing-server /http-testing-server
ENTRYPOINT ["/http-testing-server"]

EXPOSE 8080