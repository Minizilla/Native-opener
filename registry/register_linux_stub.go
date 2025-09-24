//go:build !linux

package registry

func RegisterOnLinux(protocol, progPath, args string) {
	// empty declaration, for compilation / typing
}
