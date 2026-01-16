# SVG Tool

![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-blue.svg)

A high-performance CLI tool written in Go to convert SVG files into multiple PNG formats and ICO files optimized for Web, PWA, and Mobile.
A modern, cross-platform, dependency-free (runtime) replacement for ImageMagick-based scripts.

## 🚀 Features

- **Zero Dependencies:** Does not require ImageMagick or external libraries installed on the system.
- **Cross-Platform:** Runs natively on Linux, Windows, and macOS.
- **Fast:** Processes the SVG only once in memory.
- **Native ICO:** Generates multi-layer `favicon.ico` files (16x16, 32x32, 48x48) without external dependencies.

## 🛠 Installation

### Prerequisites
You need [Go installed](https://go.dev/dl/) to compile the project. Once compiled, the binary is self-sufficient and portable.

### Linux / macOS
Clone the repository and use `make` to install:

```bash
git clone https://your-repo/svg-tool.git
cd svg-tool
sudo make install
WindowsClone the repository and execute the installation script via PowerShell:PowerShell.\scripts\install.ps1
📖 Usage1. Standard Web Kit (Recommended)Generates all necessary icons for a modern website (Favicon, Apple Touch Icon, Android PWA).Bashsvg-tool -file logo.svg -dir public/img
Output:favicon.ico (Multi-layer)favicon-16.pngfavicon-32.pngapple-touch-icon.png (180x180)pwa-android.png (192x192)2. Custom SizesGenerates specific icons, useful for browser extension manifests or UI assets.Bashsvg-tool -file logo.svg -dir assets -sizes 64,128,256
Output:icon-64.pngicon-128.pngicon-256.pngfavicon.ico (Generated from available sizes)3. Available FlagsFlagDescriptionDefault-file(Required) Path to the input SVG file.-dirOutput directory (created if it does not exist).. (Current directory)-sizesComma-separated list of sizes."" (Uses Standard Web Kit)-icoGenerates the favicon.ico file.true🧪 Development and TestingTo run the unit test suite:Bashgo test ./pkg/... -v
To build locally without installing:Bashgo build -o svg-tool cmd/svg-tool/main.go

---
