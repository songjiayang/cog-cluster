FROM golang:alpine3.20 AS builder
WORKDIR /build
COPY . .
RUN GOPROXY=off go build -o cog-agent -mod=vendor ./agent/.

FROM alpine
WORKDIR /app 
COPY --from=builder /build/cog-agent /app/cog-agent
RUN mkdir -p /app/logs
CMD ["./cog-agent"]