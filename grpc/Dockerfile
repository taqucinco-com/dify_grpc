# Dockerfile
FROM golang:1.22.1

WORKDIR /usr/src/app

COPY ./src .
EXPOSE 8080
EXPOSE 50051

RUN apt-get update
RUN apt-get install -y protobuf-compiler

# https://grpc.io/docs/languages/go/quickstart/
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
RUN go install github.com/joho/godotenv/cmd/godotenv@latest

# RUN go mod init github.com/taqucinco-com/dify_grpc
# RUN go mod download
# RUN go mod verify
# RUN go mod tidy
# RUN go build -v -o /usr/local/bin/app ./main.go
# CMD ["app"]
