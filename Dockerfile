FROM golang:latest AS base
WORKDIR /chess
COPY . .
RUN go mod download
RUN go run github.com/prisma/prisma-client-go generate

FROM base as dev
RUN go install github.com/prisma/prisma-client-go@latest
RUN go install github.com/cosmtrek/air@latest
EXPOSE 80
CMD prisma-client-go migrate deploy; air;

FROM base AS builder
RUN go build -o chess.exe ./cmd/server/server.go

FROM golang:latest
WORKDIR /chess
COPY --from=builder /chess/chess.exe ./chess.exe
COPY --from=builder /chess/prisma ./prisma
RUN go install github.com/prisma/prisma-client-go@latest
EXPOSE 80
CMD prisma-client-go migrate deploy; ./chess.exe
