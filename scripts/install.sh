#!/bin/sh
set -e

# Port Detective automated installer for Linux and macOS
# Usage: curl -sSfL https://raw.githubusercontent.com/Khoa180806/Port_Detective/master/scripts/install.sh | sh

OWNER="Khoa180806"
REPO="Port_Detective"
BINARY="pd"

echo "==> Detecting operating system and architecture..."

OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

case "$OS" in
    linux)
        OS_NAME="Linux"
        ;;
    darwin)
        OS_NAME="Darwin"
        ;;
    *)
        echo "Error: Unsupported operating system '$OS'. Port Detective supports Linux and macOS." >&2
        exit 1
        ;;
esac

case "$ARCH" in
    x86_64|amd64)
        ARCH_NAME="x86_64"
        ;;
    aarch64|arm64)
        ARCH_NAME="arm64"
        ;;
    *)
        echo "Error: Unsupported CPU architecture '$ARCH'." >&2
        exit 1
        ;;
esac

echo "==> Fetching latest release information..."
LATEST_TAG=$(curl -sSf "https://api.github.com/repos/${OWNER}/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_TAG" ]; then
    echo "Warning: Could not fetch latest release tag via GitHub API, falling back to v1.1.0"
    LATEST_TAG="v1.1.0"
fi

VERSION="${LATEST_TAG#v}"
TARBALL="${REPO}_${OS_NAME}_${ARCH_NAME}.tar.gz"
DOWNLOAD_URL="https://github.com/${OWNER}/${REPO}/releases/download/${LATEST_TAG}/${TARBALL}"

echo "==> Downloading Port Detective ${LATEST_TAG} for ${OS_NAME} (${ARCH_NAME})..."
TMP_DIR=$(mktemp -d)
trap 'rm -rf "$TMP_DIR"' EXIT

curl -sSfL "$DOWNLOAD_URL" -o "${TMP_DIR}/${TARBALL}"

echo "==> Extracting archive..."
tar -xzf "${TMP_DIR}/${TARBALL}" -C "$TMP_DIR"

# Determine target installation directory in PATH
INSTALL_DIR=""
if [ -w "/usr/local/bin" ]; then
    INSTALL_DIR="/usr/local/bin"
elif [ -d "$HOME/.local/bin" ] && echo "$PATH" | grep -q "$HOME/.local/bin"; then
    INSTALL_DIR="$HOME/.local/bin"
elif [ -w "/usr/bin" ]; then
    INSTALL_DIR="/usr/bin"
else
    INSTALL_DIR="$HOME/.local/bin"
    mkdir -p "$INSTALL_DIR"
fi

echo "==> Installing '${BINARY}' to ${INSTALL_DIR}..."
if [ -w "$INSTALL_DIR" ]; then
    mv "${TMP_DIR}/${BINARY}" "${INSTALL_DIR}/${BINARY}"
    chmod +x "${INSTALL_DIR}/${BINARY}"
else
    sudo mv "${TMP_DIR}/${BINARY}" "${INSTALL_DIR}/${BINARY}"
    sudo chmod +x "${INSTALL_DIR}/${BINARY}"
fi

echo ""
echo "========================================================="
echo "  Port Detective installed successfully!"
echo "  Run 'pd --help' or 'pd check 8080' to get started."
echo "========================================================="
