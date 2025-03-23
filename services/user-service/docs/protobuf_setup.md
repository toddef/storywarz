# Protocol Buffer Setup Guide for Windows

This guide will help you set up Protocol Buffers (protoc) on Windows for the StoryWarz user-service.

## Installing Protocol Buffer Compiler (protoc)

1. **Download the protocol buffer compiler**:
   - Go to [GitHub Releases Page](https://github.com/protocolbuffers/protobuf/releases)
   - Download the appropriate zip file for Windows (e.g., `protoc-25.3-win64.zip` for 64-bit Windows)

2. **Extract the zip file** to a location on your computer (e.g., `C:\protoc`)

3. **Add the bin directory to your PATH**:
   - Right-click on "This PC" or "My Computer" and select "Properties"
   - Click on "Advanced system settings"
   - Click on "Environment Variables"
   - Under "System Variables", find and select the "Path" variable, then click "Edit"
   - Click "New" and add the path to the `bin` directory (e.g., `C:\protoc\bin`)
   - Click "OK" on all dialog boxes

4. **Verify the installation**:
   - Open a new PowerShell window (important to get the updated PATH)
   - Run `protoc --version`
   - You should see output like `libprotoc 25.3` (version may vary)

## Installing Go Plugins for Protocol Buffers

1. **Install the Go plugin for Protocol Buffers**:
   ```
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   ```

2. **Install the Go plugin for gRPC**:
   ```
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

3. **Verify the installation**:
   - These commands should install the plugins to your `$GOPATH/bin` directory
   - Ensure this directory is in your PATH

## Generating Protocol Buffer Code

After installing the Protocol Buffer Compiler and Go plugins, you can generate the code by running:

```
cd services/user-service
.\scripts\generate_proto_windows.ps1
```

Or using the Makefile (if running in Git Bash or similar):

```
cd services/user-service
make proto
```

## Troubleshooting

If you encounter an error like `protoc: command not found`:
- Make sure you've added the Protocol Buffer Compiler's bin directory to your PATH
- Make sure you've opened a new terminal window after updating the PATH
- Try restarting your computer to ensure the PATH changes take effect

If you encounter an error like `Option --go_out: protoc-gen-go: Plugin failed with status code 1`:
- Make sure you've installed the Go plugins correctly
- Make sure your `$GOPATH/bin` directory is in your PATH
- Try reinstalling the Go plugins 