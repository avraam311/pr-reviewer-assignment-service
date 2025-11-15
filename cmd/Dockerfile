FROM golang:alpine AS build_base

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o app ./cmd/main.go

FROM alpine AS runner

WORKDIR /app

RUN mkdir -p /app/logs

COPY --from=build_base ./app .
COPY ./config ./config
COPY ./.env .

CMD ["./app"]