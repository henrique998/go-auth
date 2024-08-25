FROM golang AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o goauth ./cmd/api

FROM alpine:3.18
COPY --from=builder /app/goauth /goauth
EXPOSE 3333
CMD ["/goauth"]
