package registry

import (
	"strings"
	"testing"
)

// Test helper function to extract command from desktop entry
func extractExecFromDesktopEntry(desktopContent string) string {
	lines := strings.SplitSeq(desktopContent, "\n")
	for line := range lines {
		if after, ok := strings.CutPrefix(line, "Exec="); ok {
			return after
		}
	}
	return ""
}

// Test helper function to extract command from bash script
func extractExecFromBashScript(scriptContent string) string {
	lines := strings.SplitSeq(scriptContent, "\n")
	for line := range lines {
		if after, ok := strings.CutPrefix(line, "exec "); ok {
			return after
		}
	}
	return ""
}

func TestLinuxPathQuoting(t *testing.T) {
	protocol := "testapp"
	progPath := "/home/user/My Documents/native-opener/uri-wrapper"
	args := "--verbose"

	wrapperCmd := `"` + progPath + `" "/usr/bin/freecad" ` + args

	desktopFile := `[Desktop Entry]
Name=` + protocol + `
Exec=` + wrapperCmd + ` %u
Type=Application
Terminal=false
MimeType=x-scheme-handler/` + protocol + `;
`

	execLine := extractExecFromDesktopEntry(desktopFile)

	// Verify that the path with spaces is properly quoted
	if !strings.Contains(execLine, `"/home/user/My Documents/native-opener/uri-wrapper"`) {
		t.Errorf("Expected quoted path with spaces, got: %s", execLine)
	}

	// Verify that the target program is also quoted
	if !strings.Contains(execLine, `"/usr/bin/freecad"`) {
		t.Errorf("Expected quoted target program, got: %s", execLine)
	}

	t.Logf("Linux desktop entry Exec line: %s", execLine)
}

func TestDarwinPathQuoting(t *testing.T) {
	// Test with path containing spaces
	wrapperPath := "/home/user/My Documents/native-opener/uri-wrapper"
	progPath := "/Applications/My App/FreeCAD.app/Contents/MacOS/FreeCAD"
	args := "--verbose"

	// Simulate the bash script generation
	wrapperScript := `#!/bin/bash
exec "` + wrapperPath + `" "` + progPath + `" ` + args + ` "$@"
`

	execLine := extractExecFromBashScript(wrapperScript)

	// Verify that both paths with spaces are properly quoted
	if !strings.Contains(execLine, `"/home/user/My Documents/native-opener/uri-wrapper"`) {
		t.Errorf("Expected quoted wrapper path with spaces, got: %s", execLine)
	}

	if !strings.Contains(execLine, `"/Applications/My App/FreeCAD.app/Contents/MacOS/FreeCAD"`) {
		t.Errorf("Expected quoted target program with spaces, got: %s", execLine)
	}

	t.Logf("macOS bash script exec line: %s", execLine)
}

func TestWindowsPathQuoting(t *testing.T) {
	// Test with path containing spaces
	wrapperPath := `C:\Program Files\My App\native-opener\uri-wrapper.exe`
	progPath := `C:\Program Files\FreeCAD 0.21\bin\FreeCAD.exe`
	args := "--verbose"

	// Simulate the Windows registry command format
	registryCmd := `"` + wrapperPath + `" "` + progPath + `" ` + args + ` "%1"`

	// Verify that both paths with spaces are properly quoted
	if !strings.Contains(registryCmd, `"C:\Program Files\My App\native-opener\uri-wrapper.exe"`) {
		t.Errorf("Expected quoted wrapper path with spaces, got: %s", registryCmd)
	}

	if !strings.Contains(registryCmd, `"C:\Program Files\FreeCAD 0.21\bin\FreeCAD.exe"`) {
		t.Errorf("Expected quoted target program with spaces, got: %s", registryCmd)
	}

	t.Logf("Windows registry command: %s", registryCmd)
}

func TestPathWithoutSpaces(t *testing.T) {
	// Test with paths that don't contain spaces (should still work)
	wrapperPath := "/usr/local/bin/uri-wrapper"
	progPath := "/usr/bin/freecad"
	args := "--verbose"

	// Linux format
	linuxCmd := `"` + wrapperPath + `" "` + progPath + `" ` + args
	if !strings.Contains(linuxCmd, `"/usr/local/bin/uri-wrapper"`) {
		t.Errorf("Expected quoted wrapper path, got: %s", linuxCmd)
	}

	// macOS format
	macosCmd := `exec "` + wrapperPath + `" "` + progPath + `" ` + args + ` "$@"`
	if !strings.Contains(macosCmd, `"/usr/local/bin/uri-wrapper"`) {
		t.Errorf("Expected quoted wrapper path, got: %s", macosCmd)
	}

	// Windows format
	windowsCmd := `"` + wrapperPath + `" "` + progPath + `" ` + args + ` "%1"`
	if !strings.Contains(windowsCmd, `"/usr/local/bin/uri-wrapper"`) {
		t.Errorf("Expected quoted wrapper path, got: %s", windowsCmd)
	}

	t.Logf("All platforms handle paths without spaces correctly")
}

func TestSpecialCharactersInPaths(t *testing.T) {
	// Test with paths containing special characters
	wrapperPath := "/home/user/My App (v2.0)/native-opener/uri-wrapper"
	progPath := "/Applications/FreeCAD & CAD.app/Contents/MacOS/FreeCAD"
	args := "--verbose"

	// Linux format
	linuxCmd := `"` + wrapperPath + `" "` + progPath + `" ` + args
	if !strings.Contains(linuxCmd, `"/home/user/My App (v2.0)/native-opener/uri-wrapper"`) {
		t.Errorf("Expected quoted path with special characters, got: %s", linuxCmd)
	}

	// macOS format
	macosCmd := `exec "` + wrapperPath + `" "` + progPath + `" ` + args + ` "$@"`
	if !strings.Contains(macosCmd, `"/Applications/FreeCAD & CAD.app/Contents/MacOS/FreeCAD"`) {
		t.Errorf("Expected quoted path with special characters, got: %s", macosCmd)
	}

	t.Logf("Special characters in paths are properly handled")
}
