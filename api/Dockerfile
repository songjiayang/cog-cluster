FROM golang:alpine3.20 AS builder
WORKDIR /build
COPY . .
RUN GOPROXY=off go build -o cog-api -mod=vendor ./api/.

FROM alpine
WORKDIR /app 
COPY --from=builder /build/cog-api /app/cog-api
RUN mkdir -p /app/logs
CMD ["./cog-api"]