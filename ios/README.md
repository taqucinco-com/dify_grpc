1. brewでパッケージをインストール

```
$ brew install protobuf swift-protobuf grpc-swift
```

2. コードを生成する自作シェルを実行する

```
$ sh generate.sh
```

- *.grpc.swift
- *.pb.swift

が生成される

3. コードを書く

参考となるサンプルは以下。  
https://github.com/grpc/grpc-swift
https://github.com/apple/swift-nio
https://github.com/apple/swift-protobuf
https://swiftpackageindex.com/grpc/grpc-swift/main/tutorials/grpccore/route-guide

