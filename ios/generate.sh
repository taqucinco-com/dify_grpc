#!/bin/sh

# Homebrewのbinへのパス
BREW_BIN=/opt/homebrew/bin

protoc ./*.proto \
    --proto_path=. \
    --plugin=$BREW_BIN/protoc-gen-swift \
    --plugin=$BREW_BIN/protoc-gen-grpc-swift \
    --swift_opt=Visibility=Public \
    --grpc-swift_opt=_V2=true \
    --swift_out=. \
    --grpc-swift_out=.
