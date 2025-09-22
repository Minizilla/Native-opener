package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: microzilla <protocol> <programName>")
		os.Exit(1)
	}

	protocol := os.Args[1]
	programName := os.Args[2]

	progPath, err := filepath.Abs(programName)
	if err != nil {
		panic(err)
	}

	switch runtime.GOOS {
	case "windows":
		registerWindows(protocol, progPath)
	case "linux":
		registerLinux(protocol, progPath)
	case "darwin":
		registerMac(protocol, progPath)
	default:
		fmt.Println("OS non supporté :", runtime.GOOS)
	}
}
