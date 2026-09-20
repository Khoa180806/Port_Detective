@echo off
setlocal

echo Building Port Detective for multiple platforms...

if not exist build mkdir build

echo [1/3] Building for Windows...
set GOOS=windows
set GOARCH=amd64
go build -o build/port-detective-windows-amd64.exe

echo [2/3] Building for Linux...
set GOOS=linux
set GOARCH=amd64
go build -o build/port-detective-linux-amd64

echo [3/3] Building for macOS (Darwin)...
set GOOS=darwin
set GOARCH=amd64
go build -o build/port-detective-darwin-amd64

echo.
echo Build complete. Binaries are in the "build" directory.
dir build

endlocal
