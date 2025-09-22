package microzilla

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"golang.org/x/sys/windows/registry"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: microzilla <protocol> <programName>")
		os.Exit(1)
	}

	protocol := os.Args[1]
	programName := os.Args[2]

	// Ajout de l'extension selon l'OS
	switch runtime.GOOS {
	case "windows":
		programName += ".exe"
	case "linux", "darwin":
		programName += ".bin"
	}

	// Chemin absolu du programme cible
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

func registerWindows(protocol, progPath string) {
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

	cmdKey, _, _ := registry.CreateKey(k, `shell\open\command`, registry.SET_VALUE)
	cmdKey.SetStringValue("", fmt.Sprintf("\"%s\" \"%%1\"", progPath))
	cmdKey.Close()

	fmt.Printf("✅ Protocole %s:// enregistré -> %s\n", protocol, progPath)
}

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

func registerMac(protocol, progPath string) {
	fmt.Println("⚠️ Pour macOS : ajoute dans Info.plist :")
	fmt.Printf(`
<key>CFBundleURLTypes</key>
<array>
  <dict>
    <key>CFBundleURLSchemes</key>
    <array>
      <string>%s</string>
    </array>
  </dict>
</array>
`, protocol)
	fmt.Printf("Ton binaire actuel : %s\n", progPath)
	fmt.Println("Ensuite packager en .app et exécuter lsregister.")
}
