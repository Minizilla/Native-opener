package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"microzilla/spliter"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: uri-wrapper <target_program> [additional_args] <uri>")
		os.Exit(1)
	}

	// Get the target program path
	targetProgram := os.Args[1]
	
	// Get additional arguments (if any)
	var additionalArgs []string
	var uri string
	
	// Find the URI (it should be the last argument and contain ://)
	for i := len(os.Args) - 1; i >= 2; i-- {
		if strings.Contains(os.Args[i], "://") {
			uri = os.Args[i]
			// Additional args are between target program and URI
			additionalArgs = os.Args[2:i]
			break
		}
	}
	
	// If no URI found, use the last argument
	if uri == "" && len(os.Args) > 2 {
		uri = os.Args[len(os.Args)-1]
		additionalArgs = os.Args[2 : len(os.Args)-1]
	}
	
	// Extract the arguments from the URI
	extractedArgs := spliter.ExtractArgs(uri)
	
	// Build the command arguments
	var cmdArgs []string
	cmdArgs = append(cmdArgs, additionalArgs...)
	if extractedArgs != "" {
		cmdArgs = append(cmdArgs, extractedArgs)
	}
	
	// Execute the target program
	execPath, err := filepath.Abs(targetProgram)
	if err != nil {
		fmt.Printf("Error: cannot resolve path to %s: %v\n", targetProgram, err)
		os.Exit(1)
	}
	
	// Execute the program with the extracted arguments
	cmd := exec.Command(execPath, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing %s: %v\n", execPath, err)
		os.Exit(1)
	}
}
