package spliter

import (
	"testing"
)

func TestExtractArgs(t *testing.T) {
	const uri = "native-opener://test.dwg"
	const args = "test.dwg"

	if args != ExtractArgs(uri) {
		t.Error("ExtractArgs should return '", args, "' but returned '", ExtractArgs(uri), "'")
	}
}

func TestRecognizesAccents(t *testing.T) {
	const uri = "native-opener://T%C3%A9l%C3%A9chargements/bridge.dxf"
	const args = "Téléchargements/bridge.dxf"

	if args != ExtractArgs(uri) {
		t.Error("ExtractArgs should return '", args, "' but returned '", ExtractArgs(uri), "'")
	}
}

func TestDecodesSpaces(t *testing.T) {
	const uri = "myapp://Mon%20Dossier/fichier%20test.txt"
	const args = "Mon Dossier/fichier test.txt"

	if args != ExtractArgs(uri) {
		t.Error("ExtractArgs should return '", args, "' but returned '", ExtractArgs(uri), "'")
	}
}

func TestDecodesSpecialChars(t *testing.T) {
	const uri = "myapp://path%2Fwith%2Fslashes%20and%20spaces"
	const args = "path/with/slashes and spaces"

	if args != ExtractArgs(uri) {
		t.Error("ExtractArgs should return '", args, "' but returned '", ExtractArgs(uri), "'")
	}
}

func TestHandlesNoProtocol(t *testing.T) {
	const uri = "just-a-filename.txt"
	const args = "just-a-filename.txt"

	if args != ExtractArgs(uri) {
		t.Error("ExtractArgs should return '", args, "' but returned '", ExtractArgs(uri), "'")
	}
}

func TestHandlesPathsWithSpaces(t *testing.T) {
	const uri = "freecad://My Documents/bridge.dxf"
	const args = "My Documents/bridge.dxf"

	if args != ExtractArgs(uri) {
		t.Error("ExtractArgs should return '", args, "' but returned '", ExtractArgs(uri), "'")
	}
}

func TestHandlesEncodedSpaces(t *testing.T) {
	const uri = "freecad://My%20Documents/bridge.dxf"
	const args = "My Documents/bridge.dxf"

	if args != ExtractArgs(uri) {
		t.Error("ExtractArgs should return '", args, "' but returned '", ExtractArgs(uri), "'")
	}
}

func TestHandlesComplexPaths(t *testing.T) {
	const uri = "trueview://C:/Program Files/My App (v2.0)/file.dwg"
	const args = "C:/Program Files/My App (v2.0)/file.dwg"

	if args != ExtractArgs(uri) {
		t.Error("ExtractArgs should return '", args, "' but returned '", ExtractArgs(uri), "'")
	}
}

func TestHandlesEncodedSpecialChars(t *testing.T) {
	const uri = "freecad://My%20App%20%28v2.0%29/file.dxf"
	const args = "My App (v2.0)/file.dxf"

	if args != ExtractArgs(uri) {
		t.Error("ExtractArgs should return '", args, "' but returned '", ExtractArgs(uri), "'")
	}
}
