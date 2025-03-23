@echo off
echo Generating protocol buffer code...

if not exist pkg\proto mkdir pkg\proto
protoc --go_out=. --go-grpc_out=. proto/user.proto

echo Done. 