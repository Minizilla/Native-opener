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
