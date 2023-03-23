FROM golang:latest

WORKDIR /chess

COPY . .

RUN go mod download

RUN go run github.com/prisma/prisma-client-go generate

RUN go build -o chess.exe ./cmd/server/server.go

EXPOSE 8080

CMD go run github.com/prisma/prisma-client-go migrate deploy; ./chess.exe