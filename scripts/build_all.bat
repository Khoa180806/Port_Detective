@echo off
setlocal

echo Building Port Detective (pd) for multiple platforms...

if not exist build mkdir build

echo [1/3] Building for Windows (amd64)...
set GOOS=windows
set GOARCH=amd64
go build -o build/pd-windows-amd64.exe ./cmd/pd

echo [2/3] Building for Linux (amd64)...
set GOOS=linux
set GOARCH=amd64
go build -o build/pd-linux-amd64 ./cmd/pd

echo [3/3] Building for macOS (Darwin amd64)...
set GOOS=darwin
set GOARCH=amd64
go build -o build/pd-darwin-amd64 ./cmd/pd

echo.
echo Build complete. Binaries are in the "build" directory.
dir build

endlocal
