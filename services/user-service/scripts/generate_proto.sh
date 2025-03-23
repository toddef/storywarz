#!/bin/bash
echo "Generating protocol buffer code..."

mkdir -p pkg/proto
protoc --go_out=. --go-grpc_out=. proto/user.proto

echo "Done." 