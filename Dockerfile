FROM golang:alpine AS builder
WORKDIR /app
ADD go.mod go.mod
ADD go.sum go.sum
ADD testing-server.go testing-server.go
RUN go build

FROM alpine
COPY --from=builder /app/testing-server /testing-server
ENTRYPOINT ["/testing-server"]

EXPOSE 8080