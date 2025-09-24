# Contributing to Native Opener

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or later
- Git
- GitHub account

### Installing GoReleaser

GoReleaser is required for building and releasing. Choose your preferred installation method:

#### Quick Install (Recommended)
```bash
# Using go install (requires Go 1.24+)
go install github.com/goreleaser/goreleaser/v2@latest

# Or using Homebrew (macOS/Linux)
brew install goreleaser

# Or using the bash script (works everywhere)
curl -sfL https://goreleaser.com/static/run | bash -s -- check
```

#### Other Installation Methods
For more installation options including package managers, Docker, and manual installation, see the [official GoReleaser installation guide](https://goreleaser.com/install/).

#### Verify Installation
```bash
goreleaser --version
```

### Setup for Fork/Clone

If you're working with a forked repository, you'll need to set up a GitHub token:

1. **Create a GitHub Token**

   - Go to [GitHub Settings > Tokens](https://github.com/settings/tokens/new?scopes=repo,write:packages)
   - Generate a new token with `repo` and `write:packages` scopes
   - Copy your token (you won't see it again!)

2. **Configure Your Environment**

   ```bash
   export GITHUB_TOKEN=your_token_here
   ```

   **Important:** Restart your terminal after setting the token.

3. **Start Coding**
   - Make your changes
   - Commit your work
   - Push to your fork

## 🔨 Building for Development

### Local Builds (No Release)

For development and testing, you can build binaries without creating a GitHub release:

```bash
# Build for all configured platforms
goreleaser build

# Build with snapshot version (includes timestamp)
goreleaser build --snapshot

# Build only for specific platforms
goreleaser build --snapshot --id=linux-amd64

# Clean previous builds first
goreleaser build --clean
```

### Build Output

The built binaries will be available in the `dist/` directory:
```
dist/
├── native-opener-linux-amd64/
├── native-opener-windows-amd64/
├── native-opener-darwin-amd64/
└── native-opener-darwin-arm64/
```

### Testing Your Builds

```bash
# Test the built binary
./dist/native-opener-linux-amd64/native-opener --version

# Or run with your protocol
./dist/native-opener-linux-amd64/native-opener myprotocol /usr/bin/freecad
```

## 📦 Creating a Release

### Step 1: Tag the Release

```bash
git tag -a v0.1.0 -m "Release v0.1.0: Initial release"
git push origin v0.1.0
```

### Step 2: Build and Release

```bash
goreleaser release --clean
```

This will:

- Build binaries for all supported platforms
- Create a GitHub release
- Upload assets automatically

## 🔧 Development Workflow

1. **Fork the repository**
2. **Clone your fork**

   ```bash
   git clone https://github.com/your-username/native-opener.git
   cd native-opener
   ```

3. **Create a feature branch**

   ```bash
   git checkout -b feature/your-feature-name
   ```

4. **Make your changes and test**

   ```bash
   go test ./...
   go build
   ```

5. **Build with GoReleaser (without releasing)**
   ```bash
   # Build binaries for all platforms without creating a release
   goreleaser build
   
   # Or build for specific platforms only
   goreleaser build --snapshot
   ```

6. **Commit and push**

   ```bash
   git add .
   git commit -m "Add your feature"
   git push origin feature/your-feature-name
   ```

7. **Create a Pull Request**

## 📋 Code Standards

- Follow Go conventions
- Add tests for new features
- Update documentation as needed
- Use meaningful commit messages

## 🐛 Reporting Issues

When reporting issues, please include:

- OS and version
- Go version
- Steps to reproduce
- Expected vs actual behavior

## 💡 Feature Requests

We welcome feature requests! Please:

- Check existing issues first
- Describe the use case
- Explain why it would be useful

---

**Happy coding! 🎉**

