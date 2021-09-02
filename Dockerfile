FROM golang:1.16 as builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app .

FROM debian:bullseye-slim as app
WORKDIR /app
ENTRYPOINT ["/app/app"]
EXPOSE 9999
ENV BM_HOST=0.0.0.0
ENV BM_PORT=9999
COPY --from=builder --chown=1000:1000 /go/src/app/app .
USER 1000:1000
