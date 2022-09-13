FROM golang:alpine as builder

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /src/build/dev/don ./cmd/denet_task/*

FROM alpine

WORKDIR /app

COPY --from=builder /src/build/dev/don .
COPY --from=builder /src/configs/config.yml .

ENTRYPOINT ["/app/don"]