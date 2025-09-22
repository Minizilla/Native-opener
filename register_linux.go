//go:build linux

package microzilla

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func registerLinux(protocol, progPath string) {
	desktopFile := fmt.Sprintf(`[Desktop Entry]
Name=%s
Exec=%s %%u
Type=Application
Terminal=false
MimeType=x-scheme-handler/%s;
`, protocol, progPath, protocol)

	appDir := filepath.Join(os.Getenv("HOME"), ".local/share/applications")
	os.MkdirAll(appDir, 0755)
	desktopPath := filepath.Join(appDir, protocol+".desktop")
	os.WriteFile(desktopPath, []byte(desktopFile), 0644)

	exec.Command("xdg-mime", "default", protocol+".desktop", "x-scheme-handler/"+protocol).Run()

	fmt.Printf("✅ Protocole %s:// enregistré -> %s\n", protocol, progPath)
}
