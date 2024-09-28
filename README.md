# Overview

Dify APIを叩けるgRPCのサーバーを立てることを目標とする。

1. GolangのgRPCサーバーを起動する

```
$ docker compose exec app bash
# cd myapp
# go mod init taqucinco.com/myapp # この手順は新規で作成する場合に必要
    // go.modが作成される
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    grpc/dify_grpc.proto
    // dify_grpc.pb.go, dify_grpc_grpc.pb.goが作成される
$ go mod tidy
$ go run grpc/server/server.go 
```
1. 確認方法

```
$ brew install grpcurl
$ grpcurl -plaintext localhost:50051 list
    // dify_grpc.DifyGrpcService
    // grpc.reflection.v1.ServerReflection
    // grpc.reflection.v1alpha.ServerReflection
    // サービスを表示する
$ grpcurl -plaintext localhost:50051 list dify_grpc.DifyGrpcService
    // サービスのメソッドを表示する
    // dify_grpc.DifyGrpcService.SendQuestion
$ grpcurl -plaintext -d '{"name": "taro", "question": "What are the specs of the iPhone12?"}' localhost:50051 dify_grpc.DifyGrpcService.SendQuestion
$ echo -e '{"name": "taro"}\n{"name": "hanako"}' | grpcurl -d @ -plaintext localhost:50051 greeter.Greeter.SayHelloToMany
$ echo -e '{"name": "taro"}\n{"name": "hanako"}' | grpcurl -d @ -plaintext localhost:50051 greeter.Greeter.SayChat
```