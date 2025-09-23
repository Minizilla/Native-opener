//go:build darwin

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func registerMac(protocol, progPath, args string) {
	// macOS implementation for protocol registration
	// This would typically involve creating a .app bundle or using Launch Services
	// For now, we'll use a simple approach with URL schemes

	// Create the application bundle structure
	appName := protocol + ".app"
	appPath := filepath.Join(os.Getenv("HOME"), "Applications", appName)
	contentsPath := filepath.Join(appPath, "Contents")
	macosPath := filepath.Join(contentsPath, "MacOS")

	// Create directories
	os.MkdirAll(macosPath, 0755)

	// Create Info.plist
	infoPlist := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleExecutable</key>
	<string>%s</string>
	<key>CFBundleIdentifier</key>
	<string>com.native-opener.%s</string>
	<key>CFBundleName</key>
	<string>%s</string>
	<key>CFBundleURLTypes</key>
	<array>
		<dict>
			<key>CFBundleURLName</key>
			<string>%s URL</string>
			<key>CFBundleURLSchemes</key>
			<array>
				<string>%s</string>
			</array>
		</dict>
	</array>
</dict>
</plist>`, filepath.Base(progPath), protocol, protocol, protocol, protocol)

	infoPlistPath := filepath.Join(contentsPath, "Info.plist")
	os.WriteFile(infoPlistPath, []byte(infoPlist), 0644)

	// Get the path to the uri-wrapper
	wrapperPath, err := filepath.Abs("./uri-wrapper")
	if err != nil {
		// Fallback: assume wrapper is in the same directory
		wrapperPath = "./uri-wrapper"
	}

	// Create wrapper script with arguments
	// Quote paths to handle spaces in directory names
	wrapperScript := fmt.Sprintf(`#!/bin/bash
exec "%s" "%s" %s "$@"
`, wrapperPath, progPath, args)

	executablePath := filepath.Join(macosPath, filepath.Base(progPath))
	os.WriteFile(executablePath, []byte(wrapperScript), 0755)

	// Register with Launch Services
	exec.Command("open", appPath).Run()

	fmt.Printf("✅ URI %s:// registered -> %s\n", protocol, appPath)
}
