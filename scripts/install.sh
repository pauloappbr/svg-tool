#!/bin/bash

BINARY_NAME="svg-tool"
INSTALL_DIR="/usr/local/bin"
SOURCE_PATH="cmd/svg-tool/main.go"

echo "Compiling $BINARY_NAME..."
if go build -o $BINARY_NAME $SOURCE_PATH; then
    echo "Build successful."
else
    echo "Build failed. Please verify Go installation."
    exit 1
fi

echo "Installing to $INSTALL_DIR..."

# Check if directory exists
if [ ! -d "$INSTALL_DIR" ]; then
    echo "Creating directory $INSTALL_DIR..."
    sudo mkdir -p "$INSTALL_DIR"
fi

# Move binary (sudo might be required)
if sudo mv $BINARY_NAME $INSTALL_DIR/$BINARY_NAME; then
    sudo chmod +x $INSTALL_DIR/$BINARY_NAME
    echo ""
    echo "Installation successful!"
    echo "You can now run '$BINARY_NAME' from any directory."
else
    echo "Failed to move binary. Check permissions."
    exit 1
fi