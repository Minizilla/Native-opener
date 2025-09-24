package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: native-opener <protocol> <programName> [args]")
		os.Exit(1)
	}

	protocol := os.Args[1]
	programName := os.Args[2]

	// Optional arguments for the URI handler
	args := ""
	if len(os.Args) > 3 {
		args = os.Args[3]
	}

	progPath, err := filepath.Abs(programName)
	if err != nil {
		panic(err)
	}

	switch runtime.GOOS {
	case "windows":
		RegisterOnWindows(protocol, progPath, args)
	case "linux":
		RegisterOnLinux(protocol, progPath, args)
	case "darwin":
		RegisterOnMac(protocol, progPath, args)
	default:
		fmt.Println("OS non supporté :", runtime.GOOS)
	}
}
