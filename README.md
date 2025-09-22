# Linux natif
GOOS=linux GOARCH=amd64 go build -o microzilla-linux .

# Windows
GOOS=windows GOARCH=amd64 go build -o microzilla.exe .

# macOS Intel
GOOS=darwin GOARCH=amd64 go build -o microzilla-darwin .

# macOS ARM (M1/M2)
GOOS=darwin GOARCH=arm64 go build -o microzilla-darwin-arm64 .
