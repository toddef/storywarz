Write-Output "Generating protocol buffer code..."

if (-not (Test-Path -Path "pkg\proto")) {
    New-Item -ItemType Directory -Path "pkg\proto" -Force | Out-Null
}

protoc --go_out=. --go-grpc_out=. proto/user.proto

Write-Output "Done." 