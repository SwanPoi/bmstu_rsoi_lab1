FROM golang:1.24-alpine as builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app .

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/app /app/app
COPY --from=builder /app/config/envs /app/config/envs

EXPOSE 3000

CMD ["/app/app"]


