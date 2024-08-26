FROM golang AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/api/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder ./cmd/api/main .
EXPOSE 3333
CMD ["./main"]