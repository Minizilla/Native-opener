//go:build linux

package registry

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RegisterOnLinux(protocol, progPath, args string) {
	// Get the path to the uri-wrapper
	wrapperPath, err := filepath.Abs("./uri-wrapper")
	if err != nil {
		// Fallback: assume wrapper is in the same directory
		wrapperPath = "./uri-wrapper"
	}
	fmt.Println("Wrapper path:", wrapperPath)

	// Build the command: uri-wrapper <target_program> [args] <uri>
	// Quote paths to handle spaces in directory names
	wrapperCmd := fmt.Sprintf("\"%s\" \"%s\" %s", wrapperPath, progPath, args)

	desktopFile := fmt.Sprintf(`[Desktop Entry]
Name=%s
Exec=%s %%u
Type=Application
Terminal=false
MimeType=x-scheme-handler/%s;
`, protocol, wrapperCmd, protocol)

	appDir := filepath.Join(os.Getenv("HOME"), ".local/share/applications")
	os.MkdirAll(appDir, 0755)
	desktopPath := filepath.Join(appDir, protocol+".desktop")
	os.WriteFile(desktopPath, []byte(desktopFile), 0644)

	exec.Command("xdg-mime", "default", protocol+".desktop", "x-scheme-handler/"+protocol).Run()

	fmt.Printf("✅ URI %s:// registered -> %s\n", protocol, desktopPath)
}
