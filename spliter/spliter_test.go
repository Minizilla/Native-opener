package spliter

import (
	"testing"
)

func TestExtractArgs(t *testing.T) {
	const uri = "microzilla://test.dwg"
	const args = "test.dwg"

	if args != ExtractArgs(uri) {
		t.Error("ExtractArgs should return '", args, "' but returned '", ExtractArgs(uri), "'")
	}
}

func TestRecognizesAccents(t *testing.T) {
	const uri = "microzilla://T%C3%A9l%C3%A9chargements/bridge.dxf"
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
