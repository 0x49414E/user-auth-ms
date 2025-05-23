module user_auth

go 1.22.7

toolchain go1.22.9

require (
	github.com/golang-jwt/jwt/v4 v4.5.1
	golang.org/x/crypto v0.28.0
	google.golang.org/grpc v1.68.0
	google.golang.org/protobuf v1.35.1
)

require github.com/joho/godotenv v1.5.1

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
)
