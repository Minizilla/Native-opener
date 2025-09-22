//go:build windows

package microzilla

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

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
