# SVG Tool


![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat-square&logo=go)
![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen?style=flat-square) 
![Platform](https://img.shields.io/badge/platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey?style=flat-square)

A high-performance CLI tool to convert SVG files into multiple PNG formats and ICO files optimized for Web, PWA, and Mobile.

[Features](#features) • [Installation](#installation) • [Usage](#usage) • [Contributing](#contributing) • [License](#license)

## About

SVG Tool is a modern, cross-platform, dependency-free (runtime) replacement for ImageMagick-based scripts. It processes SVG files in a single pass, generating optimized icons for web, mobile, and PWA applications without requiring external system dependencies.

Perfect for CI/CD pipelines, Docker environments, and automated asset generation workflows.

## Features

- **Zero Runtime Dependencies** — No need for ImageMagick, ImageLib, or external tools
- **Cross-Platform** — Native support for Linux, Windows, and macOS
- **High Performance** — Single-pass SVG processing with minimal memory footprint
- **Multi-layer ICO** — Generates `favicon.ico` with multiple resolutions (16×16, 32×32, 48×48)
- **Batch Processing** — Generate multiple sizes and formats in one operation
- **Portable Binary** — Self-contained, no installation requirements after compilation

## Table of Contents

- [Quick Start](#quick-start)
- [Installation](#installation)
  - [From Source](#from-source)
  - [Pre-built Binaries](#pre-built-binaries)
- [Usage](#usage)
  - [Standard Web Kit](#1-standard-web-kit-recommended)
  - [Custom Sizes](#2-custom-sizes)
  - [Available Flags](#3-available-flags)
- [Development](#development)
  - [Prerequisites](#prerequisites)
  - [Building](#building)
  - [Testing](#testing)
- [Contributing](#contributing)
- [Code of Conduct](#code-of-conduct)
- [Support](#support)
- [License](#license)

## Quick Start

```bash
# Build from source
git clone https://github.com/pauloappbr/svg-tool.git
cd svg-tool
make build

# Generate standard web icons
./svg-tool -file logo.svg -dir public/img

# Custom sizes
./svg-tool -file logo.svg -dir assets -sizes 64,128,256
```

## Installation

### Prerequisites

- **Go 1.25 or later** — [Download Go](https://go.dev/dl/)
- Once compiled, the binary is self-sufficient with zero runtime dependencies

### From Source

#### Linux / macOS

```bash
git clone https://github.com/yourusername/svg-tool.git
cd svg-tool
sudo make install
```

The binary will be installed to `/usr/local/bin/svg-tool`.

#### Windows

Clone the repository and execute the installation script via PowerShell:

```powershell
git clone https://github.com/yourusername/svg-tool.git
cd svg-tool
.\scripts\install.ps1
```

### Pre-built Binaries

Pre-built binaries are available for download from the [Releases page](https://github.com/pauloappbr/svg-tool/releases).

```bash
# Linux / macOS
wget https://github.com/pauloappbr/svg-tool/releases/download/v1.0.0/svg-tool-linux-amd64
chmod +x svg-tool-linux-amd64
sudo mv svg-tool-linux-amd64 /usr/local/bin/svg-tool

# macOS (Apple Silicon)
wget https://github.com/pauloappbr/svg-tool/releases/download/v1.0.0/svg-tool-darwin-arm64
chmod +x svg-tool-darwin-arm64
sudo mv svg-tool-darwin-arm64 /usr/local/bin/svg-tool
```

## Usage

### 1. Standard Web Kit (Recommended)

Generate all necessary icons for a modern website (Favicon, Apple Touch Icon, Android PWA):

```bash
svg-tool -file logo.svg -dir public/img
```

**Output:**
```
favicon.ico (Multi-layer)
favicon-16.png
favicon-32.png
apple-touch-icon.png (180×180)
pwa-android.png (192×192)
```

### 2. Custom Sizes

Generate specific icon sizes, useful for browser extension manifests or UI assets:

```bash
svg-tool -file logo.svg -dir assets -sizes 64,128,256
```

**Output:**
```
icon-64.png
icon-128.png
icon-256.png
favicon.ico (Generated from available sizes)
```

### 3. Available Flags

| Flag | Description | Default |
|------|-------------|---------|
| `-file` | Path to the input SVG file (required) | — |
| `-dir` | Output directory (created if it does not exist) | `.` (current directory) |
| `-sizes` | Comma-separated list of custom sizes | `""` (Uses Standard Web Kit) |
| `-ico` | Generate the `favicon.ico` file | `true` |

## Development

### Building

Build the project locally without installing:

```bash
go build -o svg-tool cmd/svg-tool/main.go
```

Or use the provided Makefile:

```bash
make build
make clean
make help
```

### Testing

Run the complete test suite:

```bash
go test ./pkg/... -v
```

Run tests with coverage:

```bash
go test ./pkg/... -v -cover
```

Generate coverage report:

```bash
go test ./pkg/... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Project Structure

```
svg-tool/
├── cmd/
│   └── svg-tool/
│       └── main.go              # CLI entry point
├── pkg/
│   └── converter/
│       ├── converter.go         # Core conversion logic
│       └── converter_test.go    # Unit tests
├── scripts/
│   ├── install.sh              # Linux/macOS installer
│   └── install.ps1             # Windows installer
├── Dockerfile                   # Docker image definition
├── Makefile                     # Build automation
├── go.mod                       # Go module definition
├── go.sum                       # Go module checksums
├── CONTRIBUTING.md              # Contribution guidelines
├── CODE_OF_CONDUCT.md           # Community code of conduct
├── LICENSE                      # Project license (MIT)
└── README.md                    # This file
```

## Contributing

We welcome contributions! Please read our [Contributing Guidelines](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

### How to Contribute

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Areas for Contribution

- Bug reports and fixes
- Performance improvements
- Documentation enhancements
- Additional output format support
- Platform-specific optimizations

## Code of Conduct

This project adheres to the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

**Be respectful, inclusive, and constructive in all interactions.**

## Support

### Getting Help

- **Documentation:** See [Usage](#usage) section above
- **Issues:** [GitHub Issues](https://github.com/pauloappbr/svg-tool/issues)
- **Discussions:** [GitHub Discussions](https://github.com/pauloappbr/svg-tool/discussions)

### Reporting Bugs

Please use [GitHub Issues](https://github.com/pauloappbr/svg-tool/issues) to report bugs. Include:

- Operating system and version
- Go version (if building from source)
- Minimal reproduction steps
- Actual vs. expected behavior
- SVG file example (if possible)

### Feature Requests

Feature requests are welcome! Please create an issue and describe:

- The desired functionality
- Use cases and benefits
- Any relevant examples

## License

This project is licensed under the **MIT License** — see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Go](https://golang.org/)
- Inspired by modern CLI best practices
- Community contributions and feedback

---

**Made with Paulo Henrique**
