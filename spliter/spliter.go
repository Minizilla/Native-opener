package spliter

import (
	"net/url"
	"regexp"
)

// ExtractArgs extracts the arguments from a URI by removing the protocol part
// and decoding URL-encoded characters
// Example: "native-opener://test.dwg" -> "test.dwg"
// Example: "myapp://T%C3%A9l%C3%A9chargements/bridge.dxf" -> "Téléchargements/bridge.dxf"
func ExtractArgs(uri string) string {
	// Regex to match protocol:// and capture everything after it
	re := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9+.-]*://(.*)$`)
	matches := re.FindStringSubmatch(uri)

	if len(matches) > 1 {
		// Decode URL-encoded characters
		decoded, err := url.QueryUnescape(matches[1])
		if err != nil {
			// If decoding fails, return the original string
			return matches[1]
		}
		return decoded
	}

	// Fallback: if no protocol found, return the original string
	return uri
}
