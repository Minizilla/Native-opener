//go:build linux

package registry

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RegisterOnLinux(protocol, progPath, args string) {

	wrapperPath, err := findWrapperPath("uri-wrapper")
	if err != nil {

		wrapperPath = "./uri-wrapper"
	}
	fmt.Println("Wrapper path:", wrapperPath)

	var wrapperCmd string
	if args != "" {
		wrapperCmd = fmt.Sprintf("\"%s\" \"%s\" %s %%u", wrapperPath, progPath, args)
	} else {
		wrapperCmd = fmt.Sprintf("\"%s\" \"%s\" %%u", wrapperPath, progPath)
	}

	desktopFile := fmt.Sprintf(`[Desktop Entry]
Name=%s
Exec=%s
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
