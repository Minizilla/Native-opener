# Contributing to Native Opener

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or later
- Git
- GitHub account

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

5. **Commit and push**

   ```bash
   git add .
   git commit -m "Add your feature"
   git push origin feature/your-feature-name
   ```

6. **Create a Pull Request**

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

