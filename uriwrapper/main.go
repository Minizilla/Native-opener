package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/Minizilla/Native-opener/spliter"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: uri-wrapper <target_program> [additional_args] <uri>")
		os.Exit(1)
	}

	targetProgram := os.Args[1]

	var additionalArgs []string
	var uri string

	for i := len(os.Args) - 1; i >= 2; i-- {
		if strings.Contains(os.Args[i], "://") {
			uri = os.Args[i]
			additionalArgs = os.Args[2:i]
			break
		}
	}

	if uri == "" && len(os.Args) > 2 {
		uri = os.Args[len(os.Args)-1]
		additionalArgs = os.Args[2 : len(os.Args)-1]
	}

	extractedArgs := spliter.ExtractArgs(uri)

	var cmdArgs []string
	cmdArgs = append(cmdArgs, additionalArgs...)
	if extractedArgs != "" {
		cmdArgs = append(cmdArgs, extractedArgs)
	}

	execPath, err := filepath.Abs(targetProgram)
	if err != nil {
		fmt.Printf("Error: cannot resolve path to %s: %v\n", targetProgram, err)
		os.Exit(1)
	}

	cmd := exec.Command(execPath, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Start the process
	if err := cmd.Start(); err != nil {
		fmt.Printf("Error starting %s: %v\n", execPath, err)
		os.Exit(1)
	}

	// Set up signal handling for cleanup
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Start a goroutine to handle cleanup when the process exits
	go func() {
		// Wait for the process to finish
		cmd.Wait()

		// Clean up the downloaded file if it exists
		if extractedArgs != "" {
			cleanupFile(extractedArgs)
		}
	}()

	// Wait for signals
	<-sigChan

	// If we receive a signal, kill the process and cleanup
	cmd.Process.Kill()
	if extractedArgs != "" {
		cleanupFile(extractedArgs)
	}
}

// cleanupFile removes the downloaded file from the system
func cleanupFile(filename string) {
	// Try to find the file in common download locations
	downloadPaths := []string{
		filepath.Join(os.Getenv("HOME"), "Downloads", filename),
		filepath.Join(os.Getenv("HOME"), "Téléchargements", filename),
		filepath.Join(os.Getenv("USERPROFILE"), "Downloads", filename), // Windows
	}

	for _, path := range downloadPaths {
		if _, err := os.Stat(path); err == nil {
			if err := os.Remove(path); err != nil {
				fmt.Printf("Warning: could not remove file %s: %v\n", path, err)
			} else {
				fmt.Printf("Cleaned up file: %s\n", path)
			}
			return
		}
	}

	// If file not found in common locations, try the filename as-is
	if _, err := os.Stat(filename); err == nil {
		if err := os.Remove(filename); err != nil {
			fmt.Printf("Warning: could not remove file %s: %v\n", filename, err)
		} else {
			fmt.Printf("Cleaned up file: %s\n", filename)
		}
	}
}
