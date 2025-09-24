package registry

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func findWrapperPath(binaryName string) (string, error) {
	fmt.Printf("🔍 Searching for uri-wrapper binary: %s\n", binaryName)

	searchPaths := []string{
		"./",
		"./dist/",
		"../dist/",
		"../../dist/",
	}

	for _, basePath := range searchPaths {
		fmt.Printf("📁 Searching in: %s\n", basePath)

		entries, err := os.ReadDir(basePath)
		if err != nil {
			fmt.Printf("❌ Cannot read directory %s: %v\n", basePath, err)
			continue
		}

		for _, entry := range entries {
			if entry.IsDir() && strings.HasPrefix(entry.Name(), "uri-wrapper_") {
				fmt.Printf("📂 Found uri-wrapper directory: %s\n", entry.Name())

				if isCorrectOS(entry.Name()) {

					binaryPath := filepath.Join(basePath, entry.Name(), binaryName)
					fmt.Printf("🔍 Checking binary: %s\n", binaryPath)

					if _, err := os.Stat(binaryPath); err == nil {
						absPath, err := filepath.Abs(binaryPath)
						if err == nil {
							fmt.Printf("✅ Found uri-wrapper: %s\n", absPath)
							return absPath, nil
						} else {
							fmt.Printf("❌ Cannot get absolute path: %v\n", err)
						}
					} else {
						fmt.Printf("❌ Binary not found: %v\n", err)
					}
				} else {
					fmt.Printf("⚠️  Skipping %s (wrong OS)\n", entry.Name())
				}
			}
		}
	}

	fmt.Printf("🔄 Fallback: trying direct path ./%s\n", binaryName)
	absPath, err := filepath.Abs("./" + binaryName)
	if err != nil {
		fmt.Printf("❌ Cannot get absolute path for fallback: %v\n", err)
		return "", fmt.Errorf("cannot resolve fallback path for %s: %v", binaryName, err)
	}

	if _, err := os.Stat(absPath); err == nil {
		fmt.Printf("✅ Found uri-wrapper (fallback): %s\n", absPath)
		return absPath, nil
	}

	fmt.Printf("❌ uri-wrapper not found in any location\n")
	return "", fmt.Errorf("uri-wrapper binary '%s' not found in any of the searched locations: %v", binaryName, searchPaths)
}

func isCorrectOS(dirName string) bool {
	currentOS := runtime.GOOS
	currentArch := runtime.GOARCH

	fmt.Printf("🔍 Checking OS compatibility: %s (current: %s/%s)\n", dirName, currentOS, currentArch)

	var osMatch bool
	switch currentOS {
	case "linux":
		osMatch = strings.Contains(dirName, "_linux_")
	case "windows":
		osMatch = strings.Contains(dirName, "_windows_")
	case "darwin":
		osMatch = strings.Contains(dirName, "_darwin_")
	default:
		osMatch = true
	}

	if !osMatch {
		fmt.Printf("❌ OS mismatch: %s (expected %s)\n", dirName, currentOS)
		return false
	}

	var archMatch bool
	switch currentArch {
	case "amd64":
		archMatch = strings.Contains(dirName, "_amd64") || strings.Contains(dirName, "_x86_64")
	case "arm64":
		archMatch = strings.Contains(dirName, "_arm64")
	case "386":
		archMatch = strings.Contains(dirName, "_386") || strings.Contains(dirName, "_i386")
	case "arm":
		archMatch = strings.Contains(dirName, "_arm")
	default:
		archMatch = true
	}

	if !archMatch {
		fmt.Printf("❌ Architecture mismatch: %s (expected %s)\n", dirName, currentArch)
		return false
	}

	fmt.Printf("✅ OS and architecture match: %s\n", dirName)
	return true
}
