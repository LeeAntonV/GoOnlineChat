FROM golang:1.21.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd

RUN go build -o /handlers/chat ./cmd/

CMD ["/app/chat"]