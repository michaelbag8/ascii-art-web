# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.22-alpine AS builder

LABEL org.opencontainers.image.title="ascii-art-web"
LABEL org.opencontainers.image.description="A web-based ASCII art generator written in Go."
LABEL org.opencontainers.image.authors="Michael BAG <michaelbag8@example.com>"
LABEL org.opencontainers.image.source="https://github.com/michaelbag8/ascii-art-web"
LABEL org.opencontainers.image.version="1.0.0"
LABEL org.opencontainers.image.licenses="MIT"

WORKDIR /src

COPY go.mod go.sum* ./
COPY *.go ./
COPY banner ./banner
COPY templates ./templates
COPY static ./static

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/ascii-art-web .

# Runtime stage
FROM scratch

LABEL org.opencontainers.image.title="ascii-art-web"
LABEL org.opencontainers.image.description="A web-based ASCII art generator written in Go."
LABEL org.opencontainers.image.authors="Michael BAG <michaelbag8@example.com>"
LABEL org.opencontainers.image.version="1.0.0"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.url="https://github.com/michaelbag8/ascii-art-web"

WORKDIR /app

COPY --from=builder /app/ascii-art-web /app/ascii-art-web
COPY --from=builder /src/banner ./banner
COPY --from=builder /src/templates ./templates
COPY --from=builder /src/static ./static

EXPOSE 8080

USER 65532:65532

ENTRYPOINT ["/app/ascii-art-web"]
