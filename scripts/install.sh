#!/bin/bash

BINARY_NAME="svg-tool"
INSTALL_DIR="/usr/local/bin"

if [ ! -f "$BINARY_NAME" ]; then
    echo "Error: Binary '$BINARY_NAME' not found."
    echo "Please run 'make build' before installing."
    exit 1
fi

echo "Installing $BINARY_NAME to $INSTALL_DIR..."

if [ ! -d "$INSTALL_DIR" ]; then
    echo "Creating directory $INSTALL_DIR..."
    mkdir -p "$INSTALL_DIR"
fi

if cp "$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"; then
    chmod +x "$INSTALL_DIR/$BINARY_NAME"
    echo "Installation successful."
    echo "Run '$BINARY_NAME -help' to get started."
else
    echo "Error: Failed to install binary. Permission denied?"
    exit 1
fi