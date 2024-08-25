FROM golang AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build  ./cmd/api

FROM alpine:3.18
COPY --from=builder ./app ./
EXPOSE 3333
CMD ["./main"]
