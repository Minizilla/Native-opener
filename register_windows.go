//go:build windows

package main

import (
	"fmt"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

func RegisterOnWindows(protocol, progPath, args string) {
	k, _, err := registry.CreateKey(registry.CLASSES_ROOT, protocol, registry.SET_VALUE)
	if err != nil {
		panic(err)
	}
	defer k.Close()

	k.SetStringValue("", "URL:"+protocol+" Protocol")
	k.SetStringValue("URL Protocol", "")

	iconKey, _, _ := registry.CreateKey(k, "DefaultIcon", registry.SET_VALUE)
	iconKey.SetStringValue("", progPath+",1")
	iconKey.Close()

	// Get the path to the uri-wrapper
	wrapperPath, err := filepath.Abs("./uri-wrapper.exe")
	if err != nil {
		// Fallback: assume wrapper is in the same directory
		wrapperPath = "./uri-wrapper.exe"
	}

	cmdKey, _, _ := registry.CreateKey(k, `shell\open\command`, registry.SET_VALUE)
	cmdKey.SetStringValue("", fmt.Sprintf("\"%s\" \"%s\" %s \"%%1\"", wrapperPath, progPath, args))
	cmdKey.Close()

	fmt.Printf("✅ URI %s:// registered -> Windows Registry\n", protocol)
}
