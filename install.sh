#!/bin/sh
set -e

REPO="electather/seer-cli"
BIN="seer-cli"
INSTALL_DIR="/usr/local/bin"

# ── Detect OS ────────────────────────────────────────────────────────────────
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
case "$OS" in
  linux)  OS="linux" ;;
  darwin) OS="darwin" ;;
  *)
    echo "Unsupported OS: $OS"
    exit 1
    ;;
esac

# ── Detect architecture ───────────────────────────────────────────────────────
ARCH=$(uname -m)
case "$ARCH" in
  x86_64)  ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *)
    echo "Unsupported architecture: $ARCH"
    exit 1
    ;;
esac

# ── Resolve latest release version ───────────────────────────────────────────
API="https://api.github.com/repos/${REPO}/releases/latest"

if command -v curl > /dev/null 2>&1; then
  RESPONSE=$(curl -fsSL "$API")
elif command -v wget > /dev/null 2>&1; then
  RESPONSE=$(wget -qO- "$API")
else
  echo "curl or wget is required"
  exit 1
fi

VERSION=$(echo "$RESPONSE" | grep '"tag_name"' | sed 's/.*"tag_name": *"\([^"]*\)".*/\1/')
if [ -z "$VERSION" ]; then
  echo "Failed to fetch latest release version"
  exit 1
fi

# GoReleaser strips the leading 'v' from archive filenames
VERSION_BARE="${VERSION#v}"

# ── Download and install ──────────────────────────────────────────────────────
ARCHIVE="${BIN}_${VERSION_BARE}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/download/${VERSION}/${ARCHIVE}"

echo "Installing ${BIN} ${VERSION} (${OS}/${ARCH})..."

TMP=$(mktemp -d)
trap 'rm -rf "$TMP"' EXIT

if command -v curl > /dev/null 2>&1; then
  curl -fsSL "$URL" -o "${TMP}/${ARCHIVE}"
else
  wget -qO "${TMP}/${ARCHIVE}" "$URL"
fi

tar -xzf "${TMP}/${ARCHIVE}" -C "$TMP"

# ── Place binary ──────────────────────────────────────────────────────────────
if [ -w "$INSTALL_DIR" ]; then
  mv "${TMP}/${BIN}" "${INSTALL_DIR}/${BIN}"
elif command -v sudo > /dev/null 2>&1; then
  sudo mv "${TMP}/${BIN}" "${INSTALL_DIR}/${BIN}"
else
  # Fall back to user-local bin
  INSTALL_DIR="$HOME/.local/bin"
  mkdir -p "$INSTALL_DIR"
  mv "${TMP}/${BIN}" "${INSTALL_DIR}/${BIN}"
  echo "Installed to ${INSTALL_DIR}/${BIN}"
  echo "Make sure ${INSTALL_DIR} is in your PATH."
fi

chmod +x "${INSTALL_DIR}/${BIN}"

echo "Done. Run: ${BIN} --help"
