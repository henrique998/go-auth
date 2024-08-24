FROM golang:1.22-alpine3.18 as base

WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go build

FROM alpine:3.18 as binary
COPY --from=base /app/main .
EXPOSE 3333
CMD [ "./main" ]