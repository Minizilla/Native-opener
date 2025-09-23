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
