//go:build !windows

package registry

func RegisterOnWindows(protocol, progPath, args string) {
	// empty declaration, for compilation / typing
}
