package spliter

import (
	"regexp"
)

// ExtractArgs extracts the arguments from a URI by removing the protocol part
// Example: "microzilla://test.dwg" -> "test.dwg"
// Example: "myapp://path/to/file?param=value" -> "path/to/file?param=value"
func ExtractArgs(uri string) string {
	// Regex to match protocol:// and capture everything after it
	re := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9+.-]*://(.*)$`)
	matches := re.FindStringSubmatch(uri)

	if len(matches) > 1 {
		return matches[1]
	}

	// Fallback: if no protocol found, return the original string
	return uri
}
