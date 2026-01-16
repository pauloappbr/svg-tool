# Contributing to svg-tool

First off, thank you for considering contributing to **svg-tool**! It's people like you that make the open-source community such an amazing place to learn, inspire, and create.

## 🤝 Code of Conduct

By participating in this project, you are expected to uphold our [Code of Conduct](CODE_OF_CONDUCT.md). Please report unacceptable behavior to the project maintainers.

## 🚀 How to Contribute

### 1. Reporting Bugs
Before creating bug reports, please check the existing issues list as you might find out that you don't need to create one. When you are creating a bug report, please include as many details as possible:
* Use a clear and descriptive title.
* Describe the exact steps to reproduce the problem.
* Provide the SVG file that caused the issue (if applicable).
* Include your Operating System and Go version.

### 2. Suggesting Enhancements
Feature requests are welcome! Please create an issue to discuss the feature before you start coding. This saves time and ensures the feature aligns with the project's goals.

### 3. Pull Requests
We follow the "Fork-and-Pull" Git workflow.

1.  **Fork** the repository on GitHub.
2.  **Clone** your fork locally.
3.  **Create a Branch** for your feature or fix:
    * `git checkout -b feature/amazing-feature`
    * `git checkout -b fix/bug-fix`
4.  **Develop** your changes.
    * Ensure your code formats correctly (use `gofmt`).
    * Run tests to ensure no regressions:
        ```bash
        make test
        ```
5.  **Commit** your changes using **Conventional Commits**.
    * We use semantic commit messages to automate releases.
    * Format: `<type>(<scope>): <subject>`
    * Examples:
        * `feat: add support for jpg output`
        * `fix(converter): resolve crash on empty svg`
        * `docs: update installation guide`
6.  **Push** to your branch.
7.  **Open a Pull Request** targeting the `main` branch.

## 🛠 Development Guide

### Prerequisites
* [Go](https://go.dev/) 1.25 or higher.
* Make (optional, but recommended).

### Local Build
To build the binary locally for testing:

```bash
make build
# or
go build -o svg-tool cmd/svg-tool/main.go
```
### Running Tests
Please ensure all tests pass before submitting your PR:

```bash
make test
# or
go test ./pkg/... -v
```

## 📝 License
By contributing, you agree that your contributions will be licensed under its MIT License.