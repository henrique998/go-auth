FROM golang:1.22.5 as base

WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go build -o goauth ./cmd/api

FROM alpine:3.18 as binary
COPY --from=base /app/main .
EXPOSE 3333
CMD [ "./goauth" ]