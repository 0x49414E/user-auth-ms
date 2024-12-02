FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o user-auth-microservice main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=builder /app/user-auth-microservice .

EXPOSE 8080

CMD ["./user-auth-microservice"]
